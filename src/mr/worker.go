package mr

import (
	"bufio"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"net/rpc"
	"os"
	"sort"
	"strings"
	"time"
)

// Map functions return a slice of KeyValue.
type KeyValue struct {
	Key   string
	Value string
}

type ByKey []KeyValue

// for sorting by key.
func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

// main/mrworker.go calls this function.
func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {
	seenCoordinator := false
	for {
		task, ok := requestTask()
		if !ok {
			if seenCoordinator {
				return
			}
			time.Sleep(500 * time.Millisecond)
			continue
		}
		seenCoordinator = true

		switch task.Type {
		case TaskMap:
			if err := runMap(task, mapf); err == nil {
				reportTask(task)
			}
		case TaskReduce:
			if err := runReduce(task, reducef); err == nil {
				reportTask(task)
			}
		case TaskWait:
			time.Sleep(500 * time.Millisecond)
		case TaskExit:
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func requestTask() (Task, bool) {
	args := GetTaskArgs{}
	reply := GetTaskReply{}
	ok := call("Coordinator.GetTask", &args, &reply)
	return reply.Task, ok
}

func reportTask(task Task) {
	args := ReportTaskArgs{
		TaskType: task.Type,
		TaskId:   task.TaskId,
	}
	reply := ReportTaskReply{}
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if call("Coordinator.ReportTask", &args, &reply) {
			return
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func runMap(task Task, mapf func(string, string) []KeyValue) error {
	file, err := os.Open(task.File)
	if err != nil {
		return err
	}
	content, err := io.ReadAll(file)
	if err != nil {
		file.Close()
		return err
	}
	file.Close()

	kva := mapf(task.File, string(content))

	type mapOut struct {
		tmp   *os.File
		buf   *bufio.Writer
		enc   *json.Encoder
		final string
	}

	outputs := make([]mapOut, task.NReduce)
	for r := 0; r < task.NReduce; r++ {
		tmp, err := os.CreateTemp(".", "mr-tmp-")
		if err != nil {
			return err
		}
		buf := bufio.NewWriter(tmp)
		outputs[r] = mapOut{
			tmp:   tmp,
			buf:   buf,
			enc:   json.NewEncoder(buf),
			final: fmt.Sprintf("mr-%d-%d", task.TaskId, r),
		}
	}

	for _, kv := range kva {
		r := ihash(kv.Key) % task.NReduce
		if err := outputs[r].enc.Encode(&kv); err != nil {
			for _, out := range outputs {
				out.tmp.Close()
				_ = os.Remove(out.tmp.Name())
			}
			return err
		}
	}

	for _, out := range outputs {
		if err := out.buf.Flush(); err != nil {
			out.tmp.Close()
			_ = os.Remove(out.tmp.Name())
			return err
		}
		if err := out.tmp.Close(); err != nil {
			_ = os.Remove(out.tmp.Name())
			return err
		}
		if err := os.Rename(out.tmp.Name(), out.final); err != nil {
			_ = os.Remove(out.tmp.Name())
			return err
		}
	}

	return nil
}

func runReduce(task Task, reducef func(string, []string) string) error {
	entries, err := os.ReadDir(".")
	if err != nil {
		return err
	}

	intermediate := []KeyValue{}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !strings.HasPrefix(name, "mr-") || strings.HasPrefix(name, "mr-out-") {
			continue
		}
		var mapId int
		var reduceId int
		if _, err := fmt.Sscanf(name, "mr-%d-%d", &mapId, &reduceId); err != nil {
			continue
		}
		if reduceId != task.TaskId {
			continue
		}
		file, err := os.Open(name)
		if err != nil {
			return err
		}
		dec := json.NewDecoder(file)
		for {
			var kv KeyValue
			if err := dec.Decode(&kv); err != nil {
				if err == io.EOF {
					break
				}
				file.Close()
				return err
			}
			intermediate = append(intermediate, kv)
		}
		file.Close()
	}

	sort.Sort(ByKey(intermediate))

	tmp, err := os.CreateTemp(".", "mr-out-tmp-")
	if err != nil {
		return err
	}
	buf := bufio.NewWriter(tmp)
	i := 0
	for i < len(intermediate) {
		j := i + 1
		for j < len(intermediate) && intermediate[j].Key == intermediate[i].Key {
			j++
		}
		values := make([]string, 0, j-i)
		for k := i; k < j; k++ {
			values = append(values, intermediate[k].Value)
		}
		output := reducef(intermediate[i].Key, values)
		fmt.Fprintf(buf, "%v %v\n", intermediate[i].Key, output)
		i = j
	}
	if err := buf.Flush(); err != nil {
		_ = tmp.Close()
		_ = os.Remove(tmp.Name())
		return err
	}
	if err := tmp.Close(); err != nil {
		_ = os.Remove(tmp.Name())
		return err
	}
	finalName := fmt.Sprintf("mr-out-%d", task.TaskId)
	if err := os.Rename(tmp.Name(), finalName); err != nil {
		_ = os.Remove(tmp.Name())
		return err
	}

	return nil
}

// example function to show how to make an RPC call to the coordinator.
//
// the RPC argument and reply types are defined in rpc.go.
func CallExample() {

	// declare an argument structure.
	args := ExampleArgs{}

	// fill in the argument(s).
	args.X = 99

	// declare a reply structure.
	reply := ExampleReply{}

	// send the RPC request, wait for the response.
	// the "Coordinator.Example" tells the
	// receiving server that we'd like to call
	// the Example() method of struct Coordinator.
	ok := call("Coordinator.Example", &args, &reply)
	if ok {
		// reply.Y should be 100.
		fmt.Printf("reply.Y %v\n", reply.Y)
	} else {
		fmt.Printf("call failed!\n")
	}
}

// send an RPC request to the coordinator, wait for the response.
// usually returns true.
// returns false if something goes wrong.
func call(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	sockname := coordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		return false
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	log.Println(err)
	return false
}

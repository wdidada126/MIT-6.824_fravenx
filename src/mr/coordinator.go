package mr

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
	"time"
)

type TaskState int

const (
	TaskIdle TaskState = iota
	TaskInProgress
	TaskCompleted
)

type TaskMeta struct {
	state     TaskState
	startTime time.Time
}

type Coordinator struct {
	mu         sync.Mutex
	files      []string
	nReduce    int
	mapMeta    []TaskMeta
	reduceMeta []TaskMeta
}

// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

func (c *Coordinator) GetTask(args *GetTaskArgs, reply *GetTaskReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.resetExpired()

	if !c.allMapsDone() {
		for i := range c.mapMeta {
			if c.mapMeta[i].state == TaskIdle {
				c.mapMeta[i].state = TaskInProgress
				c.mapMeta[i].startTime = time.Now()
				reply.Task = Task{
					Type:    TaskMap,
					File:    c.files[i],
					TaskId:  i,
					NReduce: c.nReduce,
					NMap:    len(c.files),
				}
				return nil
			}
		}
		reply.Task = Task{Type: TaskWait}
		return nil
	}

	if !c.allReducesDone() {
		for i := range c.reduceMeta {
			if c.reduceMeta[i].state == TaskIdle {
				c.reduceMeta[i].state = TaskInProgress
				c.reduceMeta[i].startTime = time.Now()
				reply.Task = Task{
					Type:    TaskReduce,
					TaskId:  i,
					NReduce: c.nReduce,
					NMap:    len(c.files),
				}
				return nil
			}
		}
		reply.Task = Task{Type: TaskWait}
		return nil
	}

	reply.Task = Task{Type: TaskExit}
	return nil
}

func (c *Coordinator) ReportTask(args *ReportTaskArgs, reply *ReportTaskReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	switch args.TaskType {
	case TaskMap:
		if args.TaskId >= 0 && args.TaskId < len(c.mapMeta) {
			if c.mapMeta[args.TaskId].state != TaskCompleted {
				c.mapMeta[args.TaskId].state = TaskCompleted
			}
		}
	case TaskReduce:
		if args.TaskId >= 0 && args.TaskId < len(c.reduceMeta) {
			if c.reduceMeta[args.TaskId].state != TaskCompleted {
				c.reduceMeta[args.TaskId].state = TaskCompleted
			}
		}
	}
	return nil
}

// start a thread that listens for RPCs from worker.go
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

func (c *Coordinator) resetExpired() {
	now := time.Now()
	for i := range c.mapMeta {
		if c.mapMeta[i].state == TaskInProgress && now.Sub(c.mapMeta[i].startTime) > 10*time.Second {
			c.mapMeta[i].state = TaskIdle
		}
	}
	for i := range c.reduceMeta {
		if c.reduceMeta[i].state == TaskInProgress && now.Sub(c.reduceMeta[i].startTime) > 10*time.Second {
			c.reduceMeta[i].state = TaskIdle
		}
	}
}

func (c *Coordinator) allMapsDone() bool {
	for i := range c.mapMeta {
		if c.mapMeta[i].state != TaskCompleted {
			return false
		}
	}
	return true
}

func (c *Coordinator) allReducesDone() bool {
	for i := range c.reduceMeta {
		if c.reduceMeta[i].state != TaskCompleted {
			return false
		}
	}
	return true
}

// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
func (c *Coordinator) Done() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.allMapsDone() && c.allReducesDone()
}

// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{
		files:      files,
		nReduce:    nReduce,
		mapMeta:    make([]TaskMeta, len(files)),
		reduceMeta: make([]TaskMeta, nReduce),
	}
	c.server()
	return &c
}

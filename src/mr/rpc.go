package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import "os"
import "strconv"

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

// TaskType indicates what kind of work the coordinator is assigning.
type TaskType int

const (
	TaskMap TaskType = iota
	TaskReduce
	TaskWait
	TaskExit
)

// Task describes a unit of work for a worker.
type Task struct {
	Type    TaskType
	File    string
	TaskId  int
	NReduce int
	NMap    int
}

type GetTaskArgs struct{}

type GetTaskReply struct {
	Task Task
}

type ReportTaskArgs struct {
	TaskType TaskType
	TaskId   int
}

type ReportTaskReply struct{}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/5840-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}

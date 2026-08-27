package raft

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Debugging
const debug = 1

type logTopic string

const (
	dClient  logTopic = "CLNT"
	dCommit  logTopic = "CMIT"
	dDrop    logTopic = "DROP"
	dError   logTopic = "ERRO"
	dInfo    logTopic = "INFO"
	dLeader  logTopic = "LEAD"
	dLog     logTopic = "LOG1"
	dLog2    logTopic = "LOG2"
	dPersist logTopic = "PERS"
	dSnap    logTopic = "SNAP"
	dTerm    logTopic = "TERM"
	dTest    logTopic = "TEST"
	dTimer   logTopic = "TIMR"
	dTrace   logTopic = "TRCE"
	dVote    logTopic = "VOTE"
	dWarn    logTopic = "WARN"
)

func getVerbosity() int {
	v := os.Getenv("VERBOSE")
	level := 0
	if v != "" {
		var err error
		level, err = strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Invalid verbosity %v", v)
		}
	}
	return level
}

var debugStart time.Time
var debugVerbosity int
var debugTopics map[logTopic]bool

func getTopics() map[logTopic]bool {
	topics := make(map[logTopic]bool)
	for _, value := range strings.Split(os.Getenv("RAFT_LOG_TOPICS"), ",") {
		value = strings.TrimSpace(value)
		if value != "" {
			topics[logTopic(strings.ToUpper(value))] = true
		}
	}
	return topics
}

func init() {
	debugVerbosity = getVerbosity()
	debugTopics = getTopics()
	debugStart = time.Now()

	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}

func Debug(topic logTopic, format string, a ...interface{}) {
	if debugVerbosity >= 1 && (len(debugTopics) == 0 || debugTopics[topic]) {
		time := time.Since(debugStart).Microseconds()
		time /= 100
		prefix := fmt.Sprintf("%06d %v ", time, string(topic))
		format = prefix + format
		log.Printf(format, a...)
	}
}

func summarizeCommand(command interface{}) string {
	const maxRunes = 80
	value := []rune(fmt.Sprintf("%v", command))
	if len(value) <= maxRunes {
		return string(value)
	}
	return fmt.Sprintf("%s… (%d chars)", string(value[:maxRunes]), len(value))
}

func electionTime() int {
	return ELECTIONTIMEOUT + (rand.Int() % 300)
}

func cmpCommand(e1, e2 Entry) bool {
	return e1.Index == e2.Index && e1.Term == e2.Term
}

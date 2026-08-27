package shardkv

import (
	"fmt"
	"log"
	"os"
	"sort"
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
	dConfig  logTopic = "CONF"
	dMigrate logTopic = "MIGR"
	dKVOp    logTopic = "KVOP"
	dDedup   logTopic = "DUPL"
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
	for _, value := range strings.Split(os.Getenv("SHARDKV_LOG_TOPICS"), ",") {
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

func summarizeValue(value string) string {
	const maxRunes = 64
	runes := []rune(value)
	if len(runes) <= maxRunes {
		return value
	}
	return fmt.Sprintf("%s… (%d chars)", string(runes[:maxRunes]), len(runes))
}

func shardStateSummary(states map[int]int) string {
	working := make([]int, 0)
	missing := make([]int, 0)
	adding := make([]int, 0)
	for shard, state := range states {
		switch state {
		case WORKING:
			working = append(working, shard)
		case MISSING:
			missing = append(missing, shard)
		case ADDING:
			adding = append(adding, shard)
		}
	}
	sort.Ints(working)
	sort.Ints(missing)
	sort.Ints(adding)
	return fmt.Sprintf("working=%v missing=%v adding=%v", working, missing, adding)
}

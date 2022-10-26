package raft

import (
	"log"
)

type LogLevel string

const (
	offLogLevel   LogLevel = "Off"
	fatalLogLevel LogLevel = "Fatal"
	errorLogLevel LogLevel = "Error"
	infoLogLevel  LogLevel = "Info"
	debugLogLevel LogLevel = "Debug"
)

var logPriorities = map[LogLevel]int{
	offLogLevel:   0,
	fatalLogLevel: 1,
	errorLogLevel: 2,
	infoLogLevel:  3,
	debugLogLevel: 4,
}

type Flow string

const (
	FollowerFlow       Flow = "Follower"
	CandidateFlow      Flow = "Candidate"
	LeaderFlow         Flow = "Leader"
	LeaderElectionFlow Flow = "Election"
	LogReplicationFlow Flow = "LogReplicate"
	RebootFlow         Flow = "Reboot"
	LogCompactionFlow  Flow = "Compact"
)

type Logger struct {
	replicateID     int
	visibleFlows    map[Flow]bool
	visibleLogLevel LogLevel
}

func (l Logger) log(level LogLevel, flow Flow, message string) {
	if logPriorities[level] > logPriorities[l.visibleLogLevel] {
		return
	}

	_, ok := l.visibleFlows[flow]
	if !ok {
		return
	}

	log.Printf("[%v][%v][%v] %v\n", level, l.replicateID, flow, message)
}

func NewLogger(visibleLevel LogLevel, visibleFlows map[Flow]bool, replicateID int) Logger {
	return Logger{
		visibleFlows:    visibleFlows,
		visibleLogLevel: visibleLevel,
		replicateID:     replicateID,
	}
}
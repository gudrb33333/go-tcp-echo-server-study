package utils

import (
	"fmt"
	"os"
)

const (
	LOG_LEVEL_TRACE = 0
	LOG_LEVEL_DEBUG = 1
	LOG_LEVEL_INFO  = 2
	LOG_LEVEL_WARN  = 3
	LOG_LEVEL_ERROR = 4
	LOG_LEVEL_FATAL = 5
)

var _logLevel int

var logLevelStr = [6]string{"trace", "debug", "info", "warn", "error", "fatal"}

var outPutLog = func(level int, userID string, sessionUID uint64, msg string) {
	if level < _logLevel {
		return
	}

	fmt.Fprintf(os.Stdout, "[ %s ] %s\n", logLevelStr[level], msg)
}

func InitLog(loglevel int, logFunc func(int, string, uint64, string)) {
	_logLevel = loglevel

	if logFunc != nil {
		outPutLog = logFunc
	}
}

func LogTrace(userID string, sessionUID uint64, msg string) {
	outPutLog(LOG_LEVEL_TRACE, userID, sessionUID, msg)
}
func LogDebug(userID string, sessionUID uint64, msg string) {
	outPutLog(LOG_LEVEL_DEBUG, userID, sessionUID, msg)
}
func LogInfo(userID string, sessionUID uint64, msg string) {
	outPutLog(LOG_LEVEL_INFO, userID, sessionUID, msg)
}
func LogError(userID string, sessionUID uint64, msg string) {
	outPutLog(LOG_LEVEL_ERROR, userID, sessionUID, msg)
}

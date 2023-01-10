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

var logLevelStr = [6]string{"trace", "debug", "info", "warn", "error", "fatal"}

var (
	OutPutLog = _emptyExportLog
)

func InitLog(loglevel int, logFunc func(int, string, uint64, string)) {
	_logLevel = loglevel

	if logFunc != nil {
		OutPutLog = logFunc
	}
}

func logTrace(userID string, sessionUID uint64, msg string) {
	OutPutLog(LOG_LEVEL_TRACE, userID, sessionUID, msg)
}
func logDebug(userID string, sessionUID uint64, msg string) {
	OutPutLog(LOG_LEVEL_DEBUG, userID, sessionUID, msg)
}
func LogInfo(userID string, sessionUID uint64, msg string) {
	OutPutLog(LOG_LEVEL_INFO, userID, sessionUID, msg)
}
func logError(userID string, sessionUID uint64, msg string) {
	OutPutLog(LOG_LEVEL_ERROR, userID, sessionUID, msg)
}

// 비공개 함수
func _emptyExportLog(level int, userID string, sessionUID uint64, msg string) {
	if level < _logLevel {
		return
	}

	fmt.Fprintf(os.Stdout, "[ %s ] %s\n", logLevelStr[level], msg)
}

var _logLevel int

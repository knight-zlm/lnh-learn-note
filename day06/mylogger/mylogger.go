package mylogger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// LogLevel 日志类型
type LogLevel uint16

const (
	INVALID LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
)

// Logger Logger
type Logger interface {
	Debug(string)
	Info(string)
	Warning(string)
	Error(string)
}

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "error":
		return ERROR, nil
	default:
		err := errors.New("invalid level")
		return INVALID, err
	}
}

func getIndex(n int) (string, string, int) {
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Println(ok)
	}
	fileName := path.Base(file)
	funcName := runtime.FuncForPC(pc).Name()
	return fileName, funcName, line
}

func makeLog(lv LogLevel, msg string) string {
	var strLv string
	switch lv {
	case DEBUG:
		strLv = "DEBUG"
	case INFO:
		strLv = "INFO"
	case WARNING:
		strLv = "WARNING"
	case ERROR:
		strLv = "ERROR"
	default:
		strLv = "UNKOWN"
	}
	now := time.Now()
	fileName, funcName, lineNo := getIndex(4)
	return fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05.000"), strLv, fileName, funcName, lineNo, msg)
}

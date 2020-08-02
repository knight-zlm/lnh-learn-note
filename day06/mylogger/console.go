package mylogger

import "fmt"

// ConsoleLogger 日志对象
type ConsoleLogger struct {
	Level LogLevel
}

// NewConsoleLogger init func
func NewConsoleLogger(s string) Logger {
	le, _ := parseLogLevel(s)
	return ConsoleLogger{
		Level: le,
	}
}

func (l ConsoleLogger) enable(le LogLevel) bool {
	return le >= l.Level
}

func consoleLog(lv LogLevel, msg string) {
	data := makeLog(lv, msg)
	fmt.Printf(data)
}

// Debug init
func (l ConsoleLogger) Debug(msg string) {
	if l.enable(DEBUG) {
		consoleLog(DEBUG, msg)
	}
}

// Info Info
func (l ConsoleLogger) Info(msg string) {
	if l.enable(INFO) {
		consoleLog(INFO, msg)
	}
}

// Warning Warning
func (l ConsoleLogger) Warning(msg string) {
	if l.enable(WARNING) {
		consoleLog(WARNING, msg)
	}
}

// Error Error
func (l ConsoleLogger) Error(msg string) {
	if l.enable(ERROR) {
		consoleLog(ERROR, msg)
	}
}

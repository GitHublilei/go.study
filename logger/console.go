package logger

import (
	"fmt"
	"time"
)

// 日志向终端中输出

// ConsoleLogger 日志终端结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewConsoleLogger ConsoleLogger构造函数
func NewConsoleLogger(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
}

func log(lv LogLevel, msg string) {
	now := time.Now()
	funcName, fileName, line := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, line, msg)
}

// Debug console debug function
func (c ConsoleLogger) Debug(msg string) {
	if c.enable(DEBUG) {
		log(DEBUG, msg)
	}
}

// Info console info function
func (c ConsoleLogger) Info(msg string) {
	if c.enable(INFO) {
		log(INFO, msg)
	}
}

// Warning console warning function
func (c ConsoleLogger) Warning(msg string) {
	if c.enable(WARNING) {
		log(WARNING, msg)
	}
}

// Error console error function
func (c ConsoleLogger) Error(msg string) {
	if c.enable(ERROR) {
		log(ERROR, msg)
	}
}

// Fatal console fatal function
func (c ConsoleLogger) Fatal(msg string) {
	if c.enable(FATAL) {
		log(FATAL, msg)
	}
}

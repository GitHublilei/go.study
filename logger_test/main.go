package main

import (
	"time"

	"github.com/go.study/logger"
)

var log logger.Logger

// 测试日志相关内容
func main() {
	log = logger.NewConsoleLogger("warning")
	// log = logger.NewFileLogger("info", "./", "viclilei", 10*1024*1024)
	for {
		log.Debug("this is a debug log")
		log.Info("this is a info log")
		log.Warning("this is a warning log")

		id := 10001
		log.Error("this is a error log %v-%d", 111, id)
		log.Fatal("this is a fatal log")
		time.Sleep(time.Second * 2)
	}
}

package main

import (
	"time"

	"github.com/go.study/logger"
)

// 测试日志相关内容
func main() {
	log := logger.NewConsoleLogger("warning")
	for {
		log.Debug("this is a debug log")
		log.Info("this is a info log")
		log.Warning("this is a warning log")
		log.Error("this is a error log %v", 111)
		log.Fatal("this is a fatal log")
		time.Sleep(time.Second * 5)
	}
}

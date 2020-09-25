package taillog

import (
	"fmt"

	"github.com/hpcloud/tail"
)

// 专门从日志文件收集日志的模块

var (
	tailObj *tail.Tail
	logChan chan string
)

// Init 初始化
func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开（日志切割）
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 日志文件不存在不报错
		Poll:      true,                                 //
	}
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Printf("tail failed, err:%v\n", err)
		return
	}
	return
}

// ReadChan 读取日志
func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}

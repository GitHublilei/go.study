package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 向文件中写入日志

// FileLogger 文件日志结构体
type FileLogger struct {
	Level       LogLevel
	filePath    string   // 日志文件保存路径
	fileName    string   // 日志文件保存名字
	maxFileSize int64    // 日志文件最大size
	fileObj     *os.File // 日志文件对象
	errFileObj  *os.File
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}

	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	fl.initFile() // 按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v", err)
		return err
	}

	errFileObj, err := os.OpenFile(fullFileName+".err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v", err)
		return err
	}
	// 日志文件open完成
	f.fileObj = fileObj
	f.errFileObj = errFileObj

	return nil
}

// 判断文件是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return false
	}
	//如果当前文件大小大于最大值，返回true
	return fileInfo.Size() >= f.maxFileSize
}

// 切割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割日志文件
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return nil, err
	}
	nowStr := time.Now().Format("2006010215040500")
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogName := logName + ".bak" + nowStr
	file.Close()
	os.Rename(logName, newLogName)

	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err:%v\n", err)
		return nil, err
	}

	return fileObj, nil
}

// 判断是否需要记录日志
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

// 记录日志的方法
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, line := getInfo(3)
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				fmt.Printf("split file failed, err:%v\n", err)
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, line, msg)
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				newFile, err := f.splitFile(f.errFileObj)
				if err != nil {
					fmt.Printf("split file failed, err:%v\n", err)
					return
				}
				f.errFileObj = newFile
			}
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, line, msg)
		}
	}
}

// Debug file debug function
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

// Info file info function
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

// Warning file warning function
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

// Error file error function
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

// Fatal file fatal function
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

// Close 关闭日志文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

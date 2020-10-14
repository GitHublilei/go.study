package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
	mi manageInfo
)

type manageInfo struct {
	dealType string
	input    string
	output   string
	name     string
	fileType string
	times    []string
}

// 判断文件或目录是否存在
func checkFileIsExist(filename string) (exist bool) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	} else {
		exist = true
	}
	return exist
}

// Cmd 命令行调用
func Cmd(cmdStr string, params []string) (result string, err error) {
	cmd := exec.Command(cmdStr, params...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err = cmd.Start()

	if err != nil {
		fmt.Printf(", err:%v\n", err)
		return
	}
	err = cmd.Wait()
	result = out.String()
	return
}

// 格式转换
func (m *manageInfo) convertVideo() {
	defer wg.Done()
	in := m.input + m.name + "." + m.fileType
	out := m.input + m.output
	// cmdStr := fmt.Sprintf("ffmpeg -i %s -loglevel quiet -c copy -bsf:v h264_mp4toannexb -f mpegts %s", in, out)
	cmdStr := fmt.Sprintf("ffmpeg -i %s -vcodec copy -acodec copy %s", in, out)
	args := strings.Split(cmdStr, " ")
	msg, err := Cmd(args[0], args[1:])
	if err != nil {
		fmt.Printf("videoConvert failed, err:%v\n", err)
		return
	}
	fmt.Println(msg)
}

// 剪切视频
func (m *manageInfo) cutVideo() {
	defer wg.Done()
	in := mi.input + mi.name + "." + mi.fileType
	outName := mi.input + mi.name + "_." + mi.fileType
	if len(m.times) <= 2 {
		ffmepgCut(m.times[0], m.times[1], in, outName)
	} else {
		file, err := os.OpenFile("./filelist.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Printf("open filelist.txt failed, err:%v\n", err)
			return
		}
		defer file.Close()

		out := mi.input + "cut_"
		outStrArr := make([]string, len(m.times)/2)
		for i := 0; i < len(m.times)/2; i++ {
			outStr := out + strconv.Itoa(i) + "." + mi.fileType
			outStrArr[i] = outStr
			file.WriteString(fmt.Sprintf("file '%s'\n", outStr))
			ffmepgCut(m.times[i*2], m.times[i*2+1], in, outStr)
		}
		ffmepgConcat(outName)

		for _, value := range outStrArr {
			err := os.Remove(value)
			if err != nil {
				fmt.Printf("remove %s failed, err:%v\n", value, err)
			}
		}
		err = os.Remove("./filelist.txt")
		if err != nil {
			fmt.Printf("remove filelist.txt failed, err:%v\n", err)
			return
		}
	}
}

// 合并视频
func (m *manageInfo) concatVideo() {
	defer wg.Done()
	in := mi.input + mi.name + "." + mi.fileType
	out := mi.input + mi.name
	if len(m.times) <= 2 {
		out = out + "_." + mi.fileType
		ffmepgCut(m.times[0], m.times[1], in, out)
	} else {
		for i := 0; i < len(m.times)/2; i++ {

		}
	}
}

func ffmepgCut(start, end, in, out string) bool {
	cmdStr := fmt.Sprintf("ffmpeg -ss %s -to %s -accurate_seek -i %s -codec copy -avoid_negative_ts 1 %s", start, end, in, out)
	// 精确裁剪
	// cmdStr := fmt.Sprintf("ffmpeg -ss %s -to %s -i %s -c:v libx264 -c:a aac -strict experimental -b:a 98k %s", start, end, in, out)
	args := strings.Split(cmdStr, " ")
	_, err := Cmd(args[0], args[1:])
	if err != nil {
		fmt.Printf("video cut failed, err:%v\n", err)
		return false
	}
	return true
}

func ffmepgConcat(out string) bool {
	cmdStr := fmt.Sprintf("ffmpeg -f concat -safe 0 -i filelist.txt -c copy %s", out)
	args := strings.Split(cmdStr, " ")
	_, err := Cmd(args[0], args[1:])
	if err != nil {
		fmt.Printf("video cut failed, err:%v\n", err)
		return false
	}
	return true
}

func initInfo() bool {
	configPath := "./config.txt"
	exist := checkFileIsExist(configPath)
	if !exist {
		fmt.Println("config file not exist, please ")
		newFile, err := os.Create(configPath)
		if err != nil {
			fmt.Printf("create config text failed, err:%v\n", err)
			return false
		}
		newFile.Close()
		return false
	}

	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("open config file failed, err:%v\n", err)
		return false
	}

	configArr := strings.Fields(string(content))
	if len(configArr) < 3 {
		fmt.Println("config is not complete")
		return false
	}

	return setManageInfo(configArr)
}

func setManageInfo(infoArr []string) bool {
	mt := infoArr[0]
	filePath := infoArr[1]
	exist := checkFileIsExist(filePath)
	if !exist {
		fmt.Println("video file is not exist!!")
		return false
	}
	absPath, _ := filepath.Abs(filePath)
	mi.input, mi.name = filepath.Split(absPath)
	nameArr := strings.Split(mi.name, ".")
	mi.fileType = nameArr[len(nameArr)-1]
	nameArr = nameArr[:len(nameArr)-1]
	mi.name = strings.Join(nameArr, ".")
	if mt == "cut" || mt == "剪切" {
		mi.dealType = "cut"
		mi.times = infoArr[2:]
	} else if mt == "convert" || mt == "转换" {
		mi.dealType = "convert"
		mi.output = mi.name + "." + infoArr[2]
	} else if mt == "concat" || mt == "合并" {
		mi.dealType = "concat"
	}
	fmt.Printf("%v\n", mi)
	return true
}

func main() {
	start := time.Now()

	isInit := initInfo()

	if !isInit {
		fmt.Println("init config info failed")
		return
	}
	// 不知道为什么总是不行 ---------- ？？？？？？
	// refmi := reflect.ValueOf(&mi)
	// dealVideo := refmi.MethodByName("convertVideo")
	// if !dealVideo.IsValid() {
	// 	fmt.Println("not correct")
	// 	return
	// }
	// args := make([]reflect.Value, 0)
	// dealVideo.Call(args)
	wg.Add(1)
	if mi.dealType == "cut" {
		mi.cutVideo()
	} else if mi.dealType == "convert" {
		mi.convertVideo()
	} else if mi.dealType == "concat" {
		mi.concatVideo()
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("running time", elapsed)
}

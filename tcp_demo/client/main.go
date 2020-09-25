package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client

func main() {
	// 1. 与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("dial tcp failed, err:%v\n", err)
		return
	}
	// 2. 发送数据
	var msg string
	// if len(os.Args) < 2 {
	// 	msg = "hello world!"
	// } else {
	// 	msg = os.Args[1]
	// }
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("please input: ")
		text, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(text)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}

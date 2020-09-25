package main

import (
	"fmt"
	"net"
)

// tcp server端
func processConn(conn net.Conn) {
	var tmp [128]byte
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
}

func main() {
	// 1.本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("start tcp server failed, err:%v\n", err)
		return
	}
	for {
		// 2.等待别人建立链接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			return
		}
		// 3.与客户端通信
		go processConn(conn)
	}
}

package main

// UDP server

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Printf("listen UDP failed, err:%v\n", err)
		return
	}
	defer conn.Close()
	// 不需要建立连接， 直接收发数据
	var data [1021]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read from UDP failed, err:%v\n", err)
			return
		}
		fmt.Println(data[:n])
		reply := strings.ToUpper(string(data[:n]))
		// 发送数据
		conn.WriteToUDP([]byte(reply), addr)
	}

}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// UDP client

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Printf("dial UDP failed, err:%v\n", err)
		return
	}
	defer socket.Close()
	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("please input: ")
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))
		//收回复的数据
		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Printf("read reply msg failed, err:%v\n", err)
			return
		}
		fmt.Println("get reply msg: ", string(reply[:n]))
	}
}

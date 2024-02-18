package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	res, err := sendUDP("localhost:689", "Hello UDP！")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}

func sendUDP(addr, msg string) (string, error) {
	conn, _ := net.Dial("udp", addr)

	_, err := conn.Write([]byte(msg))

	bs := make([]byte, 1024)

	// 設定 UDP 連線期限
	conn.SetDeadline(time.Now().Add(3 * time.Second))
	len, err := conn.Read(bs)
	if err != nil {
		return "", err
	} else {
		return string(bs[:len]), err
	}
}

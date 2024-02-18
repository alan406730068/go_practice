package main

import (
	"fmt"
	"net"
)

func main() {
	// 與 tcp 不同，udp 要改用 net.ListenPacket() 來做監聽
	ln, err := net.ListenPacket("udp", ":689")
	defer ln.Close()
	if err != nil {
		panic("監聽 port 689 失敗")
	}
	fmt.Println("SERVER")
	for {
		buf := make([]byte, 1024)
		// ln.ReadFrom 分別回傳三個值
		// buf_len 代表讀入了幾個 byte
		// addr 代表客戶端的位址
		buf_len, addr, err := ln.ReadFrom(buf)

		if err != nil {
			continue
		}

		go func(ln net.PacketConn, addr net.Addr, buf []byte) {
			fmt.Printf("用戶端位址: %s\n收到: %s\n", addr, buf)
			ln.WriteTo([]byte("伺服器端已收到資料!\n"), addr)
		}(ln, addr, buf[:buf_len])
	}
}

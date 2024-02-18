package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("CLIENT")

	// 如果要在相同的裝置上開用戶端，可以使用 localhost，用來代表現在這台主機
	conn, err := net.Dial("tcp", "localhost:1450")
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	//發送訊息
	fmt.Fprintf(conn, "封印解除！")

	//接收伺服器回傳的訊息
	res := make([]byte, 64)
	conn.Read(res)
	fmt.Print(string(res))
}

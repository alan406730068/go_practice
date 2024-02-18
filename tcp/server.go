package main

import (
	"fmt"
	"net"
)

func main() {
	// 新增一個監聽
	// 監聽 tcp port 1450
	ln, err := net.Listen("tcp", ":1450")
	// defer 是一個特殊用法，會延遲函式執行
	// 也就是說 ln.Close() 會在 main() 即將結束時才執行
	defer ln.Close()

	if err != nil {
		panic("監聽 port 1450 失敗")
	}

	// 印出這個是 server (debug 方便)
	fmt.Println("SERVER")

	// server 是不能停止的，必需無時無刻監聽 port: 1450
	for {
		// Accept() 在沒有接到封包時會暫停，直到有接收到才會往下繼續執行
		conn, err := ln.Accept() //卡點1
		if err != nil {
			fmt.Println("ln.Accept() 失敗")
			continue
		}
		// 處理 conn
		// 因為處理 conn 時會沒辦法繼續監聽，所以要另開一條執行緒來處理，這樣若有其他用戶同時需要伺服器的服務時才不會塞車
		go func(conn net.Conn) {
			// 處理完後記得關閉 conn
			// 不然客戶端會不知道訊息傳完了沒
			defer conn.Close()
			req := make([]byte, 64)
			conn.Read(req) //卡點2

			// 回傳說已接收到並且關閉連線
			fmt.Fprintf(conn, "伺服器端回傳伺服器端已接收到 %s", string(req))

			// 印出接收到的訊息
			fmt.Println("伺服器已接收到", string(req))
		}(conn)
	}
}

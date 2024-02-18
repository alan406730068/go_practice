package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"student"
)

var client student.StudentServerClient

func main() {

	// 建立連線
	conn, err := grpc.Dial("localhost:3010", grpc.WithInsecure())
	if err != nil {
		fmt.Println("連線失敗：", err)
	}

	// 最後關閉連線
	defer conn.Close()

	// 用 proto 提供的 NewStudentServerClient，來建立 client
	client = student.NewStudentServerClient(conn)

	GetStudent(1, "A")
}

func GetStudent(studentId int64, class string) { // 透過protoc協定送回server(會送回protoc形式的包裹)
	res, err := client.GetStudentData(context.Background(),
		&student.GetStudentDataReq{
			StudentId: studentId,
			Class:     class,
		})
	if err != nil {
		log.Fatalf("GetStudentData error: %v", err)
	}
	fmt.Println(res) //執行結果
}

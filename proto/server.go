package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	// "os/exec"
	"student"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type studentServer struct {
}

func main() {
	var (
		err              error
		shutdownObserver = make(chan os.Signal, 1)
	)
	// 設定要監聽的 port
	lis, err := net.Listen("tcp", ":3010")
	if err != nil {
		panic(err)
	}

	// 使用 gRPC 的 NewServer meethod 來建立 gRPC Server
	grpcServer := grpc.NewServer()
	sv := &studentServer{}
	student.RegisterStudentServerServer(grpcServer, sv)

	// 在 gRPC 伺服器上註冊反射服務。
	reflection.Register(grpcServer)

	go func(gs *grpc.Server, c chan<- os.Signal) {
		err := gs.Serve(lis)

		if err != nil {
			shutdownObserver <- syscall.SIGINT
		}
	}(grpcServer, shutdownObserver)

	signal.Notify(shutdownObserver, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	//阻塞直到有信號傳入
	s := <-shutdownObserver
	fmt.Println(`Receive signal:`, s)

	//停止GRPC服務
	grpcServer.GracefulStop()
}

// 接受到包裹後會執行此func
func (s *studentServer) GetStudentData(ctx context.Context, in *student.GetStudentDataReq) (r *student.GetStudentDataRes, err error) {
	fmt.Println(in) //執行結果
	r = &student.GetStudentDataRes{
		StudentName:   "Alvin 2023/11/27",
		StudentHeigh:  8,
		StudentWeight: 45,
	}

	// cmd := exec.Command("mspaint")
	// cmd.Run()

	// if in.StudentId == 1 {
	// 	cmd := exec.Command("mspaint")
	// 	cmd.Run()
	// } else if in.StudentId == 2 {
	// 	cmd := exec.Command("calc")
	// 	cmd.Run()
	// }

	return r, err
}

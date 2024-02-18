package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type man struct {
	manNum int
	meter  int
	speed  int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg.Add(10)
	for i := 0; i < 10; i++ {
		var temp = &man{i + 1, 0, 0}
		go temp.run()
	}
	wg.Wait()
	// go func ()  {
	// 	for i := 0;i<100000;i++{
	// 		fmt.Print("此為匿名併發函式")
	// 	}
	// }() //此為函式，呼叫時要加()
}
func (a *man) run() {
	for a.meter < 100 {
		a.speed = rand.Intn(9) + 1
		a.meter += a.speed
		time.Sleep(1)
		//因為前幾位很早就開始跑了，加入sleep()後比較公平
	}
	fmt.Println(a.manNum)
	//panic("第一名")   //主動發起錯誤，讓所有程序停止
	wg.Done()
}

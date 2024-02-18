package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type team struct {
	teamNum int
	meter   int
	speed   int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg.Add(5)
	for i := 0; i < 5; i++ {
		var temp = &team{i + 1, 0, 0}
		go temp.firstRun()
	}
	wg.Wait()
}
func (a *team) firstRun() {
	for a.meter < 400 {
		a.speed = rand.Intn(9) + 1
		a.meter += a.speed
		time.Sleep(10)
	}
	ch := make(chan int) //通道的宣告方式
	fmt.Println(a.teamNum, "- 1")
	for i := 0; i < 3; i++ {
		var temp = &team{a.teamNum, 0, 0}
		go temp.run(ch, i) //會等待ch傳值回來後繼續執行
		<-ch               //若沒加這行不會卡住，run裡面的function(fnt)就不會執行  -- 返回通道
	} //取得channel的值，接續下位跑者
	fmt.Println(a.teamNum, "finish")
	wg.Done()
}
func (a *team) run(ch chan int, i int) {
	for a.meter < 400 {
		a.speed = rand.Intn(9) + 1
		a.meter += a.speed
		time.Sleep(10)
	}
	fmt.Println(a.teamNum, "-", i+2)
	ch <- 1 //送回channel的值
}

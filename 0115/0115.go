package main

import (
	"fmt"
	"math/rand"
	"time"
)

type poker struct {
	nums  []int
	color int
	count int
	up    int
	down  int
}

func main() {
	var chans [4]chan int
	for i := range chans {
		chans[i] = make(chan int)
	}
	for i := 0; i < 4; i++ {
		num := make([]int, 0) //0.梅花  1.方塊  2.愛心  3.黑桃
		temp := &poker{num, i, 0, 7, 7}
		go temp.selectNum(chans[i])
	}
	select {
	case c := <-chans[0]:
		fmt.Printf("梅花已結束，勝利的人是%d號玩家", c%2)
	case c := <-chans[1]:
		fmt.Printf("方塊已結束，勝利的人是%d號玩家", c%2)
	case c := <-chans[2]:
		fmt.Printf("愛心已結束，勝利的人是%d號玩家", c%2)
	case c := <-chans[3]:
		fmt.Printf("黑桃已結束，勝利的人是%d號玩家", c%2)
	}
}
func (p *poker) selectNum(ch chan int) {
	p.nums = append(p.nums, 7)
	rand.Seed(time.Now().UnixNano())
	for {
		repeat := false
		var temp int = rand.Intn(12) + 1
		for _, val := range p.nums {
			if temp == val {
				repeat = true
				break
			}
		}
		if repeat {
			continue
		}
		if temp == p.up+1 {
			p.nums = append(p.nums, temp)
			p.up = temp
		} else if temp == p.down-1 {
			p.nums = append(p.nums, temp)
			p.down = temp
		}
		p.count += 1
		if p.up == 13 || p.down == 1 {
			break
		}
	}
	ch <- p.count
}

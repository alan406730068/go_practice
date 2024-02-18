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
		go temp.run()
	}
	wg.Wait()
}
func (a *team) run() {
	var manCount int = 0
	for manCount < 4 {
		for a.meter < 400 {
			a.speed = rand.Intn(9) + 1
			a.meter += a.speed
		}
		manCount += 1
	}
	fmt.Println(a.teamNum)
	wg.Done()
}

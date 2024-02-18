package main

import (
	"fmt"
	"math/rand"
)

type lotto struct {
	num int
	min int
	max int
}

func main() {
	var min, max, num int
	fmt.Printf("請輸入值(最小值、最大值、數量)")
	fmt.Scanln(&min, &max, &num)
	lotto1 := lotto{ //lotto1 為結構變數
		num: num,
		min: min,
		max: max,
	}
	fmt.Println(lotto1.init())
}
func (l lotto) init() []int {
	var numbers []int
	map0 := make(map[int]int)
	for {
		if len(numbers) > l.num {
			break
		}
		var temp int = rand.Intn(l.max-l.min) + l.min
		//_,ok := map0[temp]
		if _, exists := map0[temp]; exists {
			continue
		}
		map0[temp] = temp
		numbers = append(numbers, temp)
	}
	return numbers
}

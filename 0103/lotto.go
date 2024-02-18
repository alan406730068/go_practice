package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

func main() {
	var min, max, num int
	fmt.Printf("請輸入值(最小值、最大值、數量)")
	fmt.Scanln(&min, &max, &num)
	fmt.Println(lotto(min, max, num))
}
func lotto(min, max, num int) []int {
	var numbers []int
	rand.Seed(time.Now().UnixNano())
	for {
		if len(numbers) > num {
			break
		}
		var temp int = rand.Intn(max-min) + min
		if slices.Contains(numbers, temp) {
			continue
		}
		numbers = append(numbers, temp)
	}
	return numbers
}

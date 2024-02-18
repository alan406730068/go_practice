package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

func main() {
	var numbers []int
	rand.Seed(time.Now().UnixNano())
	for {
		if len(numbers) > 5 {
			break
		}
		var temp int = rand.Intn(10) + 1
		if slices.Contains(numbers, temp) {
			continue
		}
		numbers = append(numbers, temp)
	}
	fmt.Println(numbers)
}

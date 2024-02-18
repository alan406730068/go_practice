package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

func main() {
	//var a interface{}
	go lotto(1, 50, 25)
	printTriangle(7)
}
func lotto(min, max, num int) {
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
	fmt.Println(numbers)
}
func printTriangle(num int) {
	for i := 0; i < num; i++ {
		if i != num-1 {
			for j := 0; j < num-i; j++ {
				fmt.Printf("%s", " ")
			}
			fmt.Printf("%s", "* ") //如果不適簍空的用for迴圈就可以結束了

			for k := 0; k < i-1; k++ {
				fmt.Printf("%s", "  ") //從第3層開始就有開始印出空白
			}
			if i != 0 {
				fmt.Printf("%s", "*") //除了第一層都要再列印一個*
			}
			fmt.Println()
		} else {
			for j := 0; j < num; j++ {
				fmt.Printf("%s", " *") //最後一行
			}
		}
	}
}

package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	//1.先設定好答案
	ans := make([]int, 4)
	//ans = append(ans, 1) //從最尾巴繼續加
	//ans = append(ans, 2)
	//ans = append(ans, 3)
	//ans = append(ans, 4)
	ans[0] = rand.Intn(8) + 1
	var flag bool
	for i := 1; i < len(ans); i++ {
		var temp = rand.Intn(9)
		for _, value := range ans { //在這foreach迴圈中，不能去修改slice裡面的值(只能去讀)
			if temp == value {
				i--
				flag = false
				break
			} else {
				flag = true
			}
		}
		if flag {
			ans[i] = temp
		}
	}
	for _, value := range ans {
		println(value)
	}
	//2.開始猜數字
	var guess [4]int
	var a, b, c, d int
	for {
		var correct, almost int
		fmt.Println("請開始你的表演")
		fmt.Scanln(&a, &b, &c, &d)
		guess[0] = a
		guess[1] = b
		guess[2] = c
		guess[3] = d
		for i := 0; i < 4; i++ {
			if guess[i] == ans[i] {
				correct += 1
			} else if slices.Contains(ans, guess[i]) {
				almost += 1
			}
		}
		if correct == 4 {
			println("恭喜答對")
			break
		} else {
			fmt.Printf("你的結果為 %d A %d B", correct, almost)
			fmt.Println()
		}
	}
}

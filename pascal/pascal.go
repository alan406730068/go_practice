package main

import (
	"fmt"
)

func main() {
	var num int
	for {
		fmt.Println("請輸入三角形的層數")
		fmt.Scanln(&num)
		if num >= 3 {
			break
		} else {
			fmt.Println("請輸入大於3的數字")
		}
	}
	//1.創建一個二維切片
	m, n := num, num+1
	pascal := make([][]int, m)
	for i := range pascal {
		pascal[i] = make([]int, n)
	}
	//2.設定2維切片裡面的值(dp)
	pascal[0][0] = 1
	pascal[1][0] = 1
	pascal[1][1] = 2
	pascal[1][2] = 1
	for i := 2; i < num; i++ {
		pascal[i][0] = 1
		for j := 1; j <= i; j++ {
			pascal[i][j] = pascal[i-1][j] + pascal[i-1][j-1]
		}
		pascal[i][i+1] = 1
	}
	fmt.Println(pascal)
	//3.畫出等邀三角形並且將值填入
	for i := 0; i < num; i++ {
		for j := num; j > i; j-- {
			fmt.Printf(" ")
		}
		for k := 0; k <= i+1; k++ {
			//fmt.Print("* ") //一般等腰三角形的寫法
			if i == 0 {
				fmt.Print(pascal[i][k])
				break
			} else {
				fmt.Print(pascal[i][k])
			}
		}
		fmt.Println()
	}
}

package main

import (
	"fmt"
)

func main() {
	printTriangle()
}
func printTriangle() {
	var num int
	fmt.Scanln(&num)
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

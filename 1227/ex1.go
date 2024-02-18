package main

import (
	"fmt"
)

func main() {
	var input, num int
	fmt.Scanln(&input, &num) //請先輸入數字(除數)  再輸入要進位的數字(被除數)
	hex(input, num)
}
func hex(a, b int) {
	if a < b {
		fmt.Print(a)
	} else {
		var temp int = a % b
		hex(a/b, b)
		fmt.Print(temp)
	}
}

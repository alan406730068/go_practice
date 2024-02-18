package main

import (
	"fmt"
)

func main() {
	var num1, num2 int       //條件若被除數等於0，則找到最大公因數a(除數)    除數和被除數會不斷的對換
	fmt.Scanln(&num1, &num2) //請先輸入數字(除數)  再輸入要進位的數字(被除數)
	gcd(num1, num2)
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

package main

import "fmt"

func main() {
	var a int32 = 1
	var b int32 = 1
	var c int32 = 2
	var num int
	fmt.Scanln(&num)
	for i := 2; i < num; i++ {
		c = a + b
		a = b
		b = c
	}
	println(c)
}

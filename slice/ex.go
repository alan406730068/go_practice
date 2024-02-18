package main

import (
	"fmt"
)

func main() {
	a := make([]int, 0, 10)
	b := append(a, 1, 2, 3)
	_ = append(a, 9, 8, 7)
	fmt.Println(b)
}

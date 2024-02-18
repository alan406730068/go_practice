package main

import (
	"fmt"
)

func main() {
	var a, b, c, max, min int32
	fmt.Scanln(&a, &b, &c)
	if a > b {
		if a > c {
			max = a
			if b > c {
				min = c
			} else {
				min = b
			}
		} else {
			max = c
			min = b
		}
	} else if b > c {
		max = b
		if a > c {
			min = c
		} else {
			min = a
		}
	} else {
		max = c
		min = a
	}
	fmt.Printf("最大值為%d,最小值為%d,", max, min)
}

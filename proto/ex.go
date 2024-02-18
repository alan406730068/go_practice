package main

import (
	"fmt"
)

func main() {
	teacher := &Teacher{Name: "Jack", Age: 32}
	fmt.Printf("%+v\n", teacher)

	tName := teacher.GetName()
	tAge := teacher.GetAge()
	fmt.Println(tName, tAge)
}

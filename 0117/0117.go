package main

import (
	"fmt"
	"regexp"

	"github.com/syhlion/simplenum"
)

func main() {
	var s = "A127071826"
	fmt.Println(simplenum.Round(1.1234, 2))
	re := regexp.MustCompile(`^[A-Z][1-2]\d{8}`) //建立表示法
	m := re.MatchString(s)                       //比較表示法
	fmt.Printf("%s is match ? %t", s, m)
}

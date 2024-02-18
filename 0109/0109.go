package main

import "fmt"

type stuff struct {
	name  string
	price int
}

func main() {
	p := stuff{"pencil", 10}  //結構變數
	var q = &stuff{"pen", 20} //指標變數
	fmt.Println(p.price)
	//p.plusPrice(&p)
	q.plusPrice() //接收者船址用法 只能用struct的指標變數
	fmt.Println(p.price)
	fmt.Println(q.price)
}
func (a *stuff) plusPrice() { //指標變數  他是屬於stuff的涵是
	a.price += 10
}

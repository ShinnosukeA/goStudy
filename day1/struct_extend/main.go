package main

import "fmt"

//模拟其他语言的继承功能
type animal struct {
	name string
}

//这个是animal的方法
func (a animal) move() {
	fmt.Println("唉，你看，动了动了", a.name)
}

type dog struct {
	feet int
	animal
}

//dog的方法
func (d dog) wang() {
	fmt.Println(d.name + "叫NMN")
}
func main() {
	d := dog{
		4,
		animal{
			"jj",
		},
	}
	fmt.Println(d)
	d.move()
	d.wang()
}

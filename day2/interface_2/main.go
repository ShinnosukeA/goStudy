package main

import "fmt"

type cat struct {
	age int
}

func (c cat) move() {
	fmt.Println("go...")
}

func (c cat) eat() {
	fmt.Println("eat....")
}

type mover interface {
	move()
}

type eater interface {
	eat()
}

//接口可以嵌套
type animal interface {
	mover
	eater
}

func main() {
	var a animal
	c := cat{1}
	//类型实现了嵌套中的所有接口的方法，也就是实现了嵌套着其他接口的这一接口
	a = c
	fmt.Println(a)
	var m mover
	m = c
	fmt.Println(m)
}

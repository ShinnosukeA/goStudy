package main

import "fmt"

type myInt int     //自定义一个类型
type yourInt = int //类型别名

func main() {
	var m myInt
	var n yourInt
	m = 100
	n = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	fmt.Println(n)
	fmt.Printf("%T\n", n)

}

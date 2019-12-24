package main

import "fmt"

func f1() func(int) int {
	//fmt.Println("hello")
	return f2
}

func f2(a int) int {
	return a
}

func f3(a func(int) int, x int) int {
	b := a(x)
	return b
}
func main() {
	fmt.Printf("f3 type is %T", f3)
}

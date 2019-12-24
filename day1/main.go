package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	a(s)
}
func test(a, b int, c ...int) (res, sub int) {
	res = a + b
	for _, v := range c {
		res += v
	}
	sub = a - b
	return res, sub
}

func a(s []int) {
	res, sub := test(1, 1, 1, 2, 3, 4, 5)
	fmt.Println(res, sub)
	for i := 0; i < 8; i++ {
		fmt.Println(i)
	}
	for i := 1; i < 9; i++ {
		fmt.Println(i)
	}
}

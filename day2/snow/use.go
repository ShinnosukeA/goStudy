package main

import (
	"fmt"
	"github.com/ShinnosukeA/goStudy/day2/calc"
)

func main() {
	add := calc.Add(1, 2)
	sub := calc.Sub(5, 1)
	mult := calc.Mult(3, 5)
	div := calc.Div(4, 2)
	fmt.Println(add)
	fmt.Println(sub)
	fmt.Println(mult)
	fmt.Println(div)
}

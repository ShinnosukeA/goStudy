package calc

import "fmt"

func init() {
	fmt.Println("init calc")
}
func Add(x, y int) int {
	return x + y
}

func Sub(a, b int) int {
	return a - b
}

func Div(a, b int) int {
	return a / b
}

func Mult(a, b int) int {
	return a * b
}

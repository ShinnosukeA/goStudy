package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 1.赋值 x=5 2 执行defer x=6  3. 返回x  返回6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //1 赋值 y=x=5 2 执行defer x++ --》x=6 3 返回 y
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 //1 赋值x=5 2 执行defer 函数中的x不是返回值的x 3 返回x=5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

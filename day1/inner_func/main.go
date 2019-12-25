package main

import "fmt"

func main() {
	defer func() {
		err := recover() // 恢复错误，使程序继续执行下去
		if err != nil {
			fmt.Println("recover A")
		}
	}()
	panic("this is a panic") //引发painc会使得程序崩溃
}

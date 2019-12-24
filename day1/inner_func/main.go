package main

import "fmt"

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover A")
		}
	}()
	panic("this is a panic")
}

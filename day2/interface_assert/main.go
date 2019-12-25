package main

import (
	"fmt"
)

type cat struct {
	name string
}

//类型断言
func f(a interface{}) {
	switch t := a.(type) {
	case int:
		fmt.Println("this is a int", t)
	case string:
		fmt.Println("this is a string", t)
	case cat:
		fmt.Println("this is a cat", t)
	case float32:
		fmt.Println("this is a float32", t)
	case *int:
		fmt.Println("this is a *int", t)
	}
}

func main() {
	f(100)
}

package main

import "fmt"

type student struct {
	name string
	age  int
	addr address
}

type school struct {
	name    string
	address //匿名嵌套的结构不冲突，可以使用匿名嵌套
}
type address struct {
	province, city string
}

func main() {
	stu := student{
		"ll",
		12,
		address{
			"sc",
			"cd",
		},
	}
	fmt.Println(stu)
	fmt.Printf("%p\n", &stu)
	fmt.Println(stu.name)
	fmt.Println(stu.addr.city)

}

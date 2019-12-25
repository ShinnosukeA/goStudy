package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	//结构体的初始化 方式1
	var xiaoming Person
	xiaoming.name = "小明"
	xiaoming.age = 10
	fmt.Println(xiaoming)
	fmt.Printf("%T\n", xiaoming)
	fmt.Println(xiaoming.age)
	//匿名结构体
	var dog struct {
		name string
		age  uint32
	}
	dog.age = 1
	dog.name = "dog"
	fmt.Printf("%T\n", dog)
	//初始化方式2
	xiaohuang := Person{
		name: "小黄",
		age:  11,
	}
	fmt.Println(xiaohuang)
	//初始化方式3
	xiaohong := Person{
		"小红",
		12,
	}
	fmt.Println(xiaohong)
}

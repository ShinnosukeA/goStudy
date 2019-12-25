package main

import "fmt"

type person struct {
	name, gender string
	age          int
}

func f(x person) {
	x.name = "aaaa"
}

func f2(x *person) {
	x.name = "ccccc" //根据内存地址找到了原来的值，可以直接修改原来的结构体的值
}

func main() {
	var p person
	p.name = "bbbb"
	p.age = 10
	p.gender = "male"
	f(p)                //传入的是p的副本，即拷贝了一份数据，修改的是拷贝的数据
	fmt.Println(p.name) //bbbb   原来的数据不变
	f2(&p)
	fmt.Printf(p.name)
}

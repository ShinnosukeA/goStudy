package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"` //在json中显示的样式
	Age  int    `jason:"age"`
}

func main() {
	p := person{
		"ll",
		12,
	}
	//json序列化
	v, err := json.Marshal(p)
	if err != nil {
		fmt.Println("failed.....")
	}
	fmt.Printf("%s\n", v)
	s := `{"name":"haha","age":12}`
	var p2 person
	err = json.Unmarshal([]byte(s), &p2) //要p2变量的值，需要传入指针，不然修改的是p2的副本，而不是p2变量本身的值
	if err != nil {
		fmt.Println("failed....")
	}
	fmt.Println(p2)
}

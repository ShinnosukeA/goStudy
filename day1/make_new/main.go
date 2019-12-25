package main

import (
	"fmt"
)

func main() {
	m1 := make(map[string]int, 1)
	a := &m1
	fmt.Printf("%x\n", a)
	m1["ll"] = 15
	m1["test"] = 1
	p := &m1
	fmt.Printf("%x\n", p)

	m2 := map[string]int{}
	s1 := []int{}
	s1 = append(s1, 12)
	fmt.Println(s1)
	fmt.Println(m2)

}

package main

import "fmt"

//
type speaker interface {
	speak()
}

type aaa struct {
}
type bbb struct {
}
type ccc struct {
}

//urbbrgr ou gay
func (a aaa) speak() {
	fmt.Println("a")
}
func (b bbb) speak() {
	fmt.Println("b")
}
func (c ccc) speak() {
	fmt.Println("c")
}

func he(s speaker) {
	s.speak()
}
func main() {
	var a = aaa{}
	b := bbb{}
	c := ccc{}
	he(a)
	he(b)
	he(c)

}

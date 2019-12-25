package main

import "fmt"

type student struct {
	name string
	age  int
}

func newStudent(name string, age int) student {
	return student{
		name,
		age,
	}

}
func main() {
	s := newStudent("a", 1)

	fmt.Println(s)
}

package main

import "fmt"

type dog struct {
	name string
}

func (d dog) jiao() {
	fmt.Println("叫NMN", d.name)
}

func main() {
	d := dog{"A"}
	d.jiao()
}

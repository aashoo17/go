package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) GiveName() {
	fmt.Println(h.name)
}

func main() {
	a := Human{}
	a.GiveName()
}

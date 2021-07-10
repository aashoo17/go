package main

import "fmt"

/*
Pointer: create, initialize, deref
pointer to structs
pointer receivers in struct
auto deref of pointerwith .
nil pointer
pointer indirection when using methods
*/

//create, initialize, deref ptr
func Ptr() {
	var a int = 10
	//create & initialize ptr
	var b *int = &a
	//deref - modify pointer underlying value
	*b = 20

	//create and then initialize
	var c *int
	c = &a

	//deref
	*c = 50

	//nil pointer - pointer doesn't point to anywhere valid in cpu
	var d *int = nil

	fmt.Println(b, c, d)
}

//todo: is it possible or not
//Human struct and passing them as pointer in methods
type Human struct {
	name string `default0:"default string"`
	age  int    `default0:"1"`
}

//pointer to struct
func StructPointer() {
	a := Human{
		name: "Something",
		age:  10,
	}
	b := &a
	//calling . on b derefs it automatically
	// -> is not available like we have in c for derefing pointer when field is requested
	fmt.Println(b.name)
	//FIXME: default0 syntax in struct is not working - check
	c := &Human{}
	//see this print result it shows that using blank literal like human{} for struct gives its fields some default values
	fmt.Println(*c)
	c.new()
	fmt.Println(*c)
}

//new() will initialize fields to some values
//removing this new() using default syntax using `` in struct
func (h *Human) new() {
	h.name = "something"
	h.age = 10
}

//change struct values to required ones
func (h *Human) changeHuman() {
	//auto deref pointer when using .
	h.age = 10
	h.name = "Something"
}

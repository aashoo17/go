package main

import "fmt"

//struct declaration
type MyStruct struct {
	name string
	age  int
}

func InitializeStruct() {
	//struct initialization using struct literal
	var a MyStruct = MyStruct{
		name: "MyStruct",
		age:  10,
	}

	//as we know string gets default value as "" and int 0 so => name: "", age: 0
	var b = MyStruct{}

	fmt.Println(a, b)
}

// methods on structs

//pass struct by value
func (a MyStruct)PassbyValue() {
	//read struct value
	var b = a.name
	var c = a.age

	fmt.Println(b, c)
}

//pass struct using pointer - called pointer receiver in go
func (a *MyStruct) PassbyPointer() {
	//modify struct using pointer
	a.name = "Some Name"
	a.age = 10
}


//TODO: can we give default values instead of 0, false, "" for struct fields
//for e.g. var a = MyStruct{} - should get default value for name: "some name" and age: 10
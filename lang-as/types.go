/*
types:
types under these
bool, integers, floats, string, char
how to use uintptr
what is the size of int in go
default values - bool, string, int = false, "", 0
type conversions - int(a)..
constants
*/

package main

import "fmt"

//creating const
const v = 10

//this const syntax works more like enums in c
//iota is like assigning 0
//and next variable will be assigned as iota + 1 here
//x = 0, y = 1, z = 2
const (
	//iota is just zero
	x int = iota
	y
	z
)

func main() {
	zeroValues()
}

func typesInGo() {
	//bool
	var a bool = true
	//integers
	var b int = 20
	var b1 uint = 20
	var b2 byte = 20 //byte alias for uint8
	//TODO: how to work with uintptr
	// var b3 uintptr = uintptr(&b)	//why this is not possible
	//floats
	var c float64 = 20
	//string
	var d = "Hello"
	//chars
	//rune is alias for int32 - so that unicode can be represented
	var e rune = 'a'

	fmt.Println(a, b, b1, b2, c, d, e)
}

func zeroValues() {
	var a int
	var b bool
	var c string

	//we are not giving any value to these but go gives
	//a = 0, b = false, c = ""
	fmt.Println(a, b, c)
}

func typeConversion() {
	var a int = 10
	var b = float64(a)
	fmt.Println(b)

	//similarly can be done for other types
}

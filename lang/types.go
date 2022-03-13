/*
Numbers:
integer:
go has opted for fixed type integers - 1, 2, 4, 8 bytes long
signed int types - int8, int16, int32, int64
unsigned int types - uint8, uint16, uint32, uint64
int, uint, uintptr - platform dependent on 32 bit 4 byte and on 64 bit 8 byte
aliases
rune - int32
byte - int8

TODO: operator precedence is simple with 5 level of precedence and left associativity how that is achieved
can work on spare time

boolean, numbers and strings can be compared using ==, !=, <, >, <=, >=
TODO: how string comparison works when string is utf-8 based

explicit conversion - go does not do any implicit conversion as it is common source of bugs in c
TODO: how to use uintptr for pointer types

default values - go gives default values to all primitive types
bool, string, int as false, "", 0 respectively

type conversions - int(a)
constants
*/

package main

import (
	"fmt"
)

//creating const
const v = 10

/*
this const syntax works more like enums in c
iota is like assigning 0
and next variable will be assigned as iota + 1 then iota + 2 and so on..
here x = 0, y = 1, z = 2
*/
const (
	//iota is just zero
	x int = iota
	y     //implicit - replace iota in expression with iota + 1
	z
)

//TODO: can multi const become any type apart from int
//so const works like any memory which can not be changed
const (
	a = ""
	b = "Hello"
	c = "World"
)

/*
numbers:
integers - int8, int16, int32, int64, uint8, uint16, uint32, uint64, int, uint, uintptr
*/
func Integers() {
	var a int8 = 10
	var b int16 = 10
	var c int32 = 10
	var d int64 = 10

	var e uint8 = 10
	var f uint16 = 10
	var g uint32 = 10
	var h uint64 = 10

	//todo: how to create uintptr from pointer types
	/*
	since ptr type does not allow arithmetic operations if in a scenario we need to do so
	we convert pointer into interger type and do all that arithmetic calculations 
	since pointers are processor dependent - uintptr is also processor dependent and can be 4 or 8 bytes
	*/
	var i uintptr = 10
	// var j uintptr = unsafe.Pointer(&i)
	var k int = 10
	var m uint = 10

	fmt.Println(a, b, c, d, e, f, g, h, i, k, m)
}

func Floats() {
	var a float32 = 10
	var b float64 = 10

	fmt.Println(a, b)
}

func Bools() {
	var a bool = true
	var b bool = false
	fmt.Println(a, b)
}

func Chars() {
	//unicode chars so rune (alias for int32) is used - no char type in go
	var a rune = 'A'
	var b rune = '\u0939'	//unicode chars
	fmt.Println(a,b)
}

/*
these 3 types get this default values implicitly if no value is given
bool - false
numbers - 0
string - ""
*/
func ZeroValues() {
	var a int
	var b bool
	var c string

	//we are not giving any value to these but go gives
	//a = 0, b = false, c = ""
	fmt.Println(a, b, c)
	//so even the struct will be built on top of these types so they get these default values for fields
}

/*
all type conversions in go are explicit
it is error to pass one type another type without explicit conversion
assignment and in function arguments
*/
func TypeConversion() {
	var a int = 10
	var b = float64(a) //converts int to float
	fmt.Println(a, b)
	//similarly can be done for other types
}

//type can be inferred by go and even var keyword can be omitted
func TypeInference() {
	a := 10 //inferred type as int
	fmt.Print(a)
}

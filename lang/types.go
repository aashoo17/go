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

TODO: how to use uintptr
TODO: what is the size of int in go

default values - go gives deffault values to all primitive types
bool, string, int = false, "", 0

type conversions - int(a)..
constants
*/

package main

import "fmt"

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
	//TODO: can we use another type inside const like string
	x int = iota
	y     //implicit replace iota in expression with iota + 1
	z
)

/*
numbers:
integers - int8, int16, int32, int64, uint8, uint16,
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

	var i uintptr = 10

	fmt.Println(a, b, c, d, e, f, g, h, i)
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
	fmt.Println(a)
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
}

/*
all type conversions in go are explicit
it is error to pass one type another type without explicit conversion
assignemnt and in function arguments
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
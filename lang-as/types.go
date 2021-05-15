/*
Numbers:
integer:
int8, int16, int32, int64
uint8, uint16, uint32, uint64
int, uint, uintptr = platform dependent on 32 bit 4 byte and on 64 bit 8 byte
aliases
rune - int32
byte - int8

TODO: operator precedence is simple with 5 level of precedence and left associativity how that is achieved can work on spare time
boolean, numbers and strings can be compared using ==, !=, <, >, <=, >=
TODO: how string comparison works when string is utf-8 based

explicit conversion

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

/*
this const syntax works more like enums in c
iota is like assigning 0
and next variable will be assigned as iota + 1 here
x = 0, y = 1, z = 2
*/
const (
	//iota is just zero
	x int = iota
	y     //implicit replace iota in expression with iota + 1
	z
)

/*
numbers:
integers - int8, int16, int32, int64, uint8, uint16,
*/
func TypesInGo() {
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
	fmt.Println(b)
	//similarly can be done for other types
}

func TypeInference() {
	a := 10 //inferred type as int
	fmt.Print(a)
}

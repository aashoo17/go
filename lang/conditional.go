package main

import "fmt"

func Conditional() {
	//if, else, else if
	var a = true
	//TODO: can we use truthy and falsy concept in go
	if a {
		fmt.Println("if branch")
	} else {
		fmt.Println("else branch")
	}

	var b = 10

	if b == 2 {
		fmt.Println("got 2")
	} else if b == 4 {
		fmt.Println("got 4")
	} else {
		fmt.Println("got other than 2 or 4")
	}
	// if with a short statement
	if a = false; a == true {
		fmt.Println("a is true then")
	}

	//switch
	var c = 20
	switch c {
	case 0:
		//do this
	case 10:
		//do this
	case 20:
		//do this
	default:
		//do if all cases fail
	}

	//switch with no condition works as - switch true
	//break are implicit unlike in c
	d := 10
	switch {
	case d == 0:
		fmt.Print(d)
	case d == 1:
		fmt.Print(d)
	case d == 2:
		fmt.Print(d)
	default:
		fmt.Print(d)
	}
}

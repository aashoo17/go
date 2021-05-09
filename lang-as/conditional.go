package main

import "fmt"

func conditional() {

	//if, else, else if
	var a = true

	if a {
		fmt.Println("if branch")
	} else {
		fmt.Println("else branch")
	}

	var b = 10

	if b == 2 {
		//do this
	} else if b == 4 {
		//do this
	} else {
		//do this
	}

	// if with a short statement
	if a = false; a == true {
		//do this
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

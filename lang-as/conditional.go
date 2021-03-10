package main

import "fmt"

func main() {

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
}

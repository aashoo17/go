package main

import "fmt"

/*
array:
create, read/modify using index
*/

func Array() {
	//create
	var a [10]int
	var b = [10]int{10, 20, 30, 40}
	c := [10]int{10, 20, 30}

	//read
	var d = b[1]

	//modify
	b[1] = 90

	fmt.Println(a, b, c, d)
}

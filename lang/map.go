package main

import "fmt"

//TODO: go map are based on hash map or btree map ?
func Map() {
	//create map
	a := map[int]int{1: 10, 2: 20}

	//update map
	a[2] = 30
	a[3] = 40
	fmt.Println(a)

	//delete a value
	delete(a, 3)
	fmt.Println(a)
}

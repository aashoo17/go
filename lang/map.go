package main

import "fmt"

//TODO: go map are based on hash map or btree map ?
func Map() {
	//create map
	a := map[int]int{1: 10, 2: 20}
	b := make(map[int]int,10)

	//update map
	a[2] = 30
	a[3] = 40
	fmt.Println(a)

	b[1] = 10
	b[2] = 20

	//delete a value
	delete(a, 3)
	fmt.Println(a,b)
}

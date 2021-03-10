package main

import "fmt"

func main() {
	//create map
	a := map[int]int{1: 10, 2: 20}

	//update map
	a[2] = 30
	a[3] = 40
	fmt.Println(a)

	//delet a value
	delete(a, 3)
	fmt.Println(a)
}

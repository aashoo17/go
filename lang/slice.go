package main

import "fmt"

/*
slice:
create from array
create from make()
modify slice - using append()
*/

func SlicWorking() {
	var a = [5]int{10, 20, 30, 40}

	//creating slice from array
	//getting slice which will point to memory inside array a
	var b []int = a[:]
	var _ = a[1:3]

	//creating slice from make()
	var c = make([]int, 10)
	fmt.Println(c)
	//changing b say will change a
	b[0] = 90
	//check if a was changed
	fmt.Println(a)
	//so arrays can not grow if we have taken reference to array using slice, growing slice can not grow array
	//so when slices are grown using append() call new memory is created somewhere
	//if append() is assigned back to slice like this b = append(b,100)
	//now changing this grown slice we can not modify the prev array

	//slice is grown and assigned back to itself so b now points to some new memory
	//changing slice b now can not change a
	//growing slice using append()
	b = append(b, 100)
	//modifying slice using index syntax
	b[0] = 70
	//print a, b and check if its true
	fmt.Println(a, b)
}

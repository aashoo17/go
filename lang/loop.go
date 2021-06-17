package main

import "fmt"

/*
in go we have only for loop no while loop
but for can become all other loop which we desire
- for loop as while
- for loop as working on iterators
- for loop running infinitely
*/

func Loop() {
	//common for syntax we are familiar with c/cpp
	//this i is block scoped
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	//no while in go - for can become while
	//since previous i was block scoped in for now it is gone we can create I again without any problem
	var i = 0
	for i < 8 {
		fmt.Print(i)
		i++
	}

	//other possible for syntaxes
	//using j here as previous i is now function scoped
	for j := 0; j < 10; {
		fmt.Print(j)
		j++
	}

	var j = 0
	for ; j < 10; j++ {
		fmt.Print(j)
	}

	//for as iterator loop
	a := []int{10, 20, 30, 40, 50}
	for k, v := range a {
		fmt.Println(k, v)
	}
}

// infinite for loop
func InfiniteLoop() {
	for {

	}
}

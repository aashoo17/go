package main

import "fmt"

func main() {
	//common for syntax we are familiar with c/cpp
	//this i is block scoped
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	//no while in go - for can become while
	//since previous i was block scoped in for now it is gone we can create i again without any problem
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
}

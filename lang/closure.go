package main

import "fmt"

//so writing closure is go is simply writing function
//but it will consume a value not available in function - as args or its local variable
//so in garbage collected language we can think it like this

//any variable created gets cleaned up when no longer in use
//but when closure references that variable garbage collector will not free it
//and keep it around for closure to be used

func Closure() {
	var a = 10
	var b = func() {
		//this closure has no variable a it is taking a variable from main function
		//function does not do that
		fmt.Println(a)
	}
	b()
	//if b would have been called somewhere generally a would have been cleaned up
	//still this b() will work printing value a as go's garbage collector keep the value around longer

	//an example of that would be if a would have returned this closure like below
	//return b
	//somewhere returned b (i.e. closure ) could be called but till that time outer function main here is gone
	//so value of a should be gone - we kinda would be accessing a memory which should be gone
	//but go keeps it around
	//TODO: using main as returning closure is not good example may be confusing for other write separate function to explain it
}

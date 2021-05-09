/*
defer:
usually go's garbage collector will cleanup resources/memory when they are not required
but some resources are heavy weight on system and should be cleaned up as soon as not required
to facilitate this go gives defer
it will be last call of the function when function finishes it task
in short words

``go
func DeferCall(){
	defer fmt.Print("I should be last line in the function to be called written in advance")
}
``

*/

package main

import "fmt"

func DeferCall() {
	defer fmt.Print("I should be last line in the function to be called written in advance")
}


/*
stacking of defers:
when multiple defers are called they get stack and called lifo (last in first out) order
*/

func MultipleDeferCall() {
	defer fmt.Print("1st call")
	defer fmt.Print("2nd call")
	defer fmt.Print("3rd call")
}
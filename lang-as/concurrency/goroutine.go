package main

import (
	"fmt"
	"time"
)

//creating a goroutine
//write a normal function
func goroutinePrint() {
	fmt.Println("Hello World")
}

func main() {
	//call the function using go like below
	//this will create a goroutine
	go goroutinePrint()
	//if main goroutine is not made sleep for 2 seconds it is possible main will exit and goroutinePrint will no run
	//2 second is not hard and fast rule it is taken as random
	//it may be possible even 2 sec is less for goroutinePrint if code inside is waiting for longer to execute or not sceduled as other goroutines are available to schedule
	//so call to time.Sleep() is not good way to block the goroutines generally and should not be used
	//TODO: find out other way than using channel
	//as if we will be using mutex generally channel will not come in picture
	//and time.Sleep() is not good tactic
	time.Sleep(2 * time.Second)
}

//now we can think of it as this
//we have created 1 goroutine named goroutinePrint and another our main goroutine = 2 goroutine in total
//main goroutine is special in a way that if it has return all other goroutine will be cleaned up even if they are not completed
//comment line time.Sleep() and see if print comes

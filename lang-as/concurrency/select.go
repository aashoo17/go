/*
in channel we can block a thread waiting for a channel to return
say we have channel ch then
<-ch
this say wait for data from channel ch

so select is kinda same with a small difference
select can block goroutine waiting for data over channel just like normal channel
but select can wait on multiple channels

//TODO: I suspect there can be more in select instead of all cases being blocking on some channel, see if it is poosible
//to write case other than being blocking on some channel
e.g.
select{
case <-ch:
	//do some
case <-ch1:
	//do some
case <-ch3:
	//do some
default:
	//run if no case channel is ready
}

so here we can wait on 3 channels using select
which one is ready first code under its case block is executed
if more than 1 ready any one will be randomly picked
TODO: will code not move past select block until 1 case is run (if default is not written, or default block is mandatory)

*/

package main

import "fmt"

func main() {
	//create 3 channels
	var ch, ch1, ch2 chan int
	//initialize the channel
	ch = make(chan int)
	ch1 = make(chan int)
	ch2 = make(chan int)

	//invoke a new goroutine
	go func(ch chan int) {
		ch <- 2
	}(ch)

	//select block
	select {
	case <-ch:
		fmt.Println("got channel ch")
	case <-ch1:
	case <-ch2:
		//TODO: if default was being given it was running instead of case <-ch:, why ??
		// default:
	}

	//TODO: what will happen if we write new goroutine here instead of above select{} block
	// go func(ch chan int) {
	// 	ch <- 2
	// }(ch)

}

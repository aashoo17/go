//channel is a way to communicate between goroutines
//passing data around - without fear the data will be modified by more than 1 goroutine together
//even blocking a goroutine if it is waiting for data over channel until data is received
//this is one way to synchronise goroutines

//1. use channel to block main goroutine so that some goroutine is executed without calling time.Sleep()
package concurrency

import "fmt"

func Channel() {
	//creating channel
	//can't use var ch chan int as the channel will be nil
	ch := make(chan int)
	go randomPrint(ch)
	//wait on the channel to receive  the int (though we are not using value, just using channel for blocking main until randomPrint() sends something on channel)
	<-ch
}

//takes a channel
func randomPrint(ch chan int) {
	fmt.Println("random print")
	//one println() is called send 10 over channel
	ch <- 10
}

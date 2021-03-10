package main

/*
buffered channel:
so they differ from normal channel as normal channel can have only 1 value
and these normal channel will block immediately as <-ch (recive from channel) this call is wriiten
buffered channel will block only if they are full
size of the channel - max value they can store before sending the data is told when they are made using make() call

*/

import "fmt"

//TODO: this example can be improved - try to do so later

func main() {
	//buffered channel of size 10
	ch := make(chan int, 10)
	go randomDataEmitter(ch)

	//since the channel can receive more than one value instead of writting <-ch (for received data)
	//they are made iterator so we can call range syntax directly
	for i := range ch {
		fmt.Println(i)
	}
}

func randomDataEmitter(ch chan int) {
	for i := 0; i <= 30; i++ {
		ch <- i
	}
	//buffered channel has to be called as closed explicitly otherwise go will not know when data is no longer being sent
	close(ch)
}

package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func WaitForGoroutines() {
	//create a waitgroup
	var wg sync.WaitGroup

	//we can add no that many times Done() has to be called on wait group
	wg.Add(1)                //wait for 1 goroutine to call Done()
	go SimulateLongTask(&wg) //pass the wait group on which Done() will be called

	wg.Wait() //wait until the goroutine calls Done()
}

func SimulateLongTask(wg *sync.WaitGroup) {
	defer wg.Done()             //call Done() on waitgroup when function is exitting
	time.Sleep(time.Second * 2) //sleep for 2 seconds
	fmt.Println("We are done")
}

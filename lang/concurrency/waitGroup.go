package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func WaitForGoroutines() {
	var wg sync.WaitGroup

	wg.Add(1)	//wait for 1 goroutine to call Done()
	go SimulateLongTask(&wg)	//pass the wait group on which Done() will be called

	wg.Wait()	//wait until the goroutine calls Done()
}

func SimulateLongTask(wg *sync.WaitGroup) {
	defer wg.Done()		//call Done() on waitgroup when function is exitting
	time.Sleep(time.Second * 2)
	fmt.Println("We are done")
}

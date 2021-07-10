package concurrency

import (
	"fmt"
	"time"
)

func goroutinePrint() {
	fmt.Println("Hello World")
}

func Goroutine() {
	go goroutinePrint()
	time.Sleep(2 * time.Second)
}


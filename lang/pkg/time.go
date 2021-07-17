package pkg

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
so Time struct keeps track of both absolute time and monotonic time in this struct - time.Time
some methods will use absolute time from epoch - calendar time
some will use monotonic time - for elapsed time
really cool api - for high level language
*/

//calendar time
func CalendarTime() {
	a := time.Now()
	fmt.Println(a) //TODO: why this is printing custom string rather than no of seconds - why go has chosen this

	//getting time, day, year etc
	//todo: we are getting time in local time zone what about GMT
	b1, b2, b3 := a.Date()
	c := a.Day()
	d1, d2, d3 := a.Clock()

	fmt.Println(b1, b2, b3)
	fmt.Println(c)
	fmt.Println(d1, d2, d3)
}

// elapsed time
func ElapsedTime() {
	a := time.Now()
	fmt.Println("Hello elapsed time")
	b := time.Now()
	c := b.Sub(a) //this uses monotonic time for substraction
	fmt.Println(c)
}

//alarm - called as Timer in go
func Alarm() {
	//create a timer which will give time after 1s duration
	b := time.NewTimer(time.Second)
	//get time over channel
	c := <-b.C
	fmt.Println(c)
}

//continuous alarm - called as Ticker in go
func ContinuousAlarm() {
	//TODO: add CTRL + C signal handling to stop the time from command line
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	//create a ticker which will give time after every 1s duration
	b := time.NewTicker(time.Second)
	go func() {
		for c := range b.C {
			fmt.Println(c)
		}
	}()
	//wait for signal - once received main thread will exit and cleanup all goroutines
	<-sigs
}

//Sleep for some time
func Sleep() {
	//wait for 2 second
	time.Sleep(time.Second * 2)
}

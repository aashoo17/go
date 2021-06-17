package pkg

import (
	"fmt"
	"time"
)

//calendar time
func CalendarTime() {
	a := time.Now()
	fmt.Println(a)
}

//Sleep for some time
func Sleep() {
	//wait for 2 second
	time.Sleep(time.Second * 2)
}

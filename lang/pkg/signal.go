package pkg

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Signal() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	//signal - registered in main thread
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		//capture sigs in this closure
		//we are waiting for signal in another thread - where signal is registered
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

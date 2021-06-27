package pkg

import (
	"os"
	"os/signal"
)

func Signal() {
	ch := make(chan os.Signal,1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}


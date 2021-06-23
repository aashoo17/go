package pkg

import (
	"fmt"
	"os/exec"
)

func CreateProcess() {
	a := exec.Command("ls", "-la")
	e := a.Run()
	if e != nil {
		fmt.Print("error")
	}

}

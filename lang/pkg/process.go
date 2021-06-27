package pkg

import (
	"fmt"
	"os"
	"os/exec"
)

func CreateProcess() {
	a := exec.Command("ls", "-la")
	a.Stdout = os.Stdout //assign the stdout to the new process by default it is nil
	b := a.Run()         //Run() wait for the completion of child process and Start() don't
	if b != nil {
		fmt.Print("error")
	}
}

func ProcessWait() {
	a := exec.Command("sleep", "3")
	a.Stdout = os.Stdout //assign the stdout to the new process by default it is nil
	b := a.Start()       //Start() does not wait for child process and child process sleeps for 3 seconds but parent process will exit
	a.Wait()             //explicit wait call is required for waiting for child, comment this line and see the difference
	if b != nil {
		fmt.Print("error")
	}
}

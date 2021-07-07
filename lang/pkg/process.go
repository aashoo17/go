package pkg

import (
	"os"
	"os/exec"
)

func CreateProcess() {
	a := exec.Command("ls", "-la")
	a.Stdout = os.Stdout //assign the stdout to the new process by default it is nil for child process
	a.Run()              //Run() wait for the completion of child process and Start() don't
}

func WaitChildProcess() {
	a := exec.Command("sleep", "3")
	a.Stdout = os.Stdout //assign the stdout to the new process by default it is nil
	a.Start()            //Start() does not wait for child process and child process sleeps for 3 seconds but parent process will exit
	a.Wait()             //explicit wait call is required for waiting for child, comment this line and see the difference
}

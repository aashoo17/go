package main

import (
	"fmt"
	"os"
	"os/exec"
)

// run a command like (ls -la)

func runCommandWaitingForChild() {
	cmd := exec.Command("ls", "-la")
	cmd.Stdout = os.Stdout
	cmd.Run() //wait for child process to complete
}

// work on global environment list
//get all env variable
func getAllEnv() {
	env := os.Environ()
	fmt.Println(env)
}

//get a particular env
func getEnv() {
	path := os.Getenv("PATH")
	fmt.Println(path)
}

func setEnv() {
	//set the env value
	//todo: does Setenv() will create the env variable if not present already
	// os.Setenv("PATH",path + ":/home/ashu/code")		//working on path
	os.Setenv("ASHU", "ashu")
}

func unsetEnv(){
	os.Unsetenv("ASHU")		//unset env variable
}

func clearEnv(){
	//clear all env variable
	os.Clearenv()
}

// todo: how to create a new env variable 
// todo: can go permanently change env variables

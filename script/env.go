package main

import (
	"fmt"
	"os"
	"os/exec"
)

// run a command like (ls -la)

func RunProcessWaitingForChild() {
	cmd := exec.Command("ls", "-la")
	cmd.Stdout = os.Stdout
	cmd.Run() //wait for child process to complete
}

/*
Environment variable:
a key value pair which are stored in particular files  
many programs/executable can change their behavior based on these key-pair value 
think of it like commandline - there we pass different flags to change executable behavior
but unlike flags environment variable can be set in advance and then no user intervention is required later
and we know that a process only creates a child process during this phase parent process will pass all of its 
environment variable to child process

all the functions used here will only modify the environment variable for current process

system wide ENVs
/etc/environment, /etc/profile, /etc/profile.d/, /etc/bash.bashrc

user wide ENVs
~/.bashrc, ~/.bash_profile, ~/.bash_login, ~/.profile

todo: I think all go functions like SetEnv, UnsetEnv, ClearEnv create the effect for current process only
we have to change the files where env are stored for permanent effect
user level envs - file starting with .bash etc are bash specific so which file in userspace to be modified
for user env changes without being specific to shell (some user may use zsh, fish..or any other) is it .profile
*/

// work on global environment list
//get all env variable
func PrintAllEnv() {
	env := os.Environ()
	fmt.Println(env)
}

//get a particular env
func GetEnv() {
	path := os.Getenv("PATH")
	fmt.Println(path)
}

func SetEnv() {
	//set the env value
	//todo: does Setenv() will create the env variable if not present already
	// os.Setenv("PATH",path + ":/home/ashu/code")		//working on path
	os.Setenv("ASHU", "ashu")
	//ASHU does gets created but will it persist - or we will have to open env storing files and write to it
	PrintAllEnv()
}

func UnsetEnv() {
	os.Unsetenv("ASHU") //unset env variable
}

func ClearEnv() {
	//clear all env variable
	os.Clearenv()
}

// todo: how to create a new env variable
// todo: can go permanently change env variables

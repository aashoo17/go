/*
Mutex:
mutex create, lock, unlock
once lock is taken on any mutex it will not allow a piece of to be run by another goroutine until unlock is called

Lock()
some code here
..some more code
Unlock()
coe between Lock() and Unlock() can be accessed by only one goorotine at once (goroutine which has taken lock only that one)
once Unlock() is called now others can have access

so idea is that if a piece of code modified from different goroutine can create data/race or race condition
keep that piece of code between Lock() and Unlock()

*/

package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func Mutex_working() {
	//create map
	m := map[int]int{1: 10}
	//TODO: initialization of mutext is as easy as calling Mutex{}, we did not require its fields updation how
	//is it possible if we don't have constructor concept in go
	mux := sync.Mutex{}
	//TODO: what will happen if we have not taken pointer type of map and mutex
	//is everything in go is copy type when passed in function
	go criticalPieceOfCode(&m, &mux)
	go criticalPieceOfCode(&m, &mux)

	//wait main thread for few seconds so goroutines can be executed
	time.Sleep(time.Second)
	fmt.Println(m)
}

//this is the code whose some part should not be accessed by more than 1 gorutine
func criticalPieceOfCode(m *map[int]int, mux *sync.Mutex) {
	//do somethong non critical here - may be print
	fmt.Println("non critical code")

	//say we are modifying map now this is not thread safe - updation of go map from multiple thread should be synchronised as lib does not provide thhat
	//only alllow 1 goroutine at a time
	//take a lock over mutex
	mux.Lock()
	//defer the unlock() which means this will be called when func criticalPieceOfCode() returns
	defer mux.Unlock()
	//update some indexes
	//TODO: try to get the current size of map and then update next 2 indices of map - so that updation from more than 1 goroutine
	//can be visible in map
	//here both goroutine are updating index 2 and 3 so not much effect is visible
	(*m)[2] = 20
	(*m)[3] = 30
	//if we don't want use defer can call unlock here
	//mux.Unlock() 	//like this, for our uses there is no code after (*m)[3] = 30 so it will not matter much in explicit Unlock() call vs in defer
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4) //2 goroutines

	go func() {
		say("Goroutines 1!") //A goroutine is a lightweight thread managed by the Go runtime.
		waitGroup.Done()     //-1 to the WaitGroup counter.
	}()

	go func() {
		say("Goroutines 2!")
		waitGroup.Done() //-1 to the WaitGroup counter.
	}()

	go func() {
		say("Goroutines 3!") //A goroutine is a lightweight thread managed by the Go runtime.
		waitGroup.Done()     //-1 to the WaitGroup counter.
	}()

	go func() {
		say("Goroutines 4!")
		waitGroup.Done() //-1 to the WaitGroup counter.
	}()
	waitGroup.Wait() //Blocks until the WaitGroup counter is 0.
}

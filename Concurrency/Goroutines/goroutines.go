package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world") //A goroutine is a lightweight thread managed by the Go runtime.
	say("hello")
}

// func main() {
// 	go escrever("Hello World!") //goroutine
// 	escrever("Programando em Go!")
// }

// func escrever(texto string) {
// 	for {
// 		fmt.Println(texto)
// 		time.Sleep(time.Second)
// 	}
// }

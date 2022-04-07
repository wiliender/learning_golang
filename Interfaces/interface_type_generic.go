package main

import "fmt"

func generica(iterf interface{}) {
	fmt.Printf("%T\n", iterf)
}

func main() {
	generica(3)
	generica("Wiliender")
	generica(true)
}

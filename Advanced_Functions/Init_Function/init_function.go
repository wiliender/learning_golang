package main

import "fmt"

var n int

func init() {
	fmt.Println("Função init sendo executada") //run before main function
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}

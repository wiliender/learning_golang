package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Printf("%s %d\n", texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo: ", 1, 2, 3, 4, 5, 6, 7, 123, 1245, 235, 12)
}

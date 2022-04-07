package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 30
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}

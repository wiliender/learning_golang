package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) { // multiple returns
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	soma := somar(10, 23)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Hello World")
	fmt.Println(resultado)

	resultadoSoma, resultadosub := calculosMatematicos(10, 23)
	fmt.Println(resultadoSoma, resultadosub)

	resultadoSoma2, _ := calculosMatematicos(10, 23) // _ is a parameter that will not be used
	fmt.Println(resultadoSoma2)
}

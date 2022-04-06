package main

import "fmt"

func main() {

	/* func() {
		fmt.Println("Olá Mundo!")
	}() */

	//POR PARAMETRO
	/* func(texto string) {
		fmt.Println(texto)
	}("Passando Parametro") */

	//POR RETORNO
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, 123)
	}("Passando Parametro")

	fmt.Println(retorno)
}

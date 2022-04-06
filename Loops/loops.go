package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	//i := 0

	/* for i < 10 {
		i++
		fmt.Println("Incrementando i:", i)
		time.Sleep(time.Second)
	}
	fmt.Println(i) */

	/* for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j:", j)
		time.Sleep(time.Second)
	} */
	//fmt.Println(j) j is not visible here

	//FOR COM CLAUSULA RANGE
	nomes := [3]string{"João", "Maria", "José"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nomes := range nomes {
		fmt.Println(nomes)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Loop infinito")
		time.Sleep(time.Second)
	}
}

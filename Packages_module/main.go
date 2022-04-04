package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo a partir do pacote main")
	auxiliar.Escrever()

	fmt.Println(checkmail.ValidateFormat("wiliendersilva"))
	erro := (checkmail.ValidateFormat("wiliendersilva@gmail.com"))
	fmt.Println(erro)
}

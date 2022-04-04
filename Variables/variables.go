package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2" //inferencia de tipo refente ao valor dele
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	const contante1 string = "Contante 1"
	fmt.Println(contante1)

	variavel5, variavel6 = variavel6, variavel5 //troca de valores sem uso de uma variável auxiliar
	fmt.Println(variavel5, variavel6)
}

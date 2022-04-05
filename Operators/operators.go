package main

import "fmt"

func main() {
	//ARITMETICOS
	soma := 1 + 2
	subtrcao := 10 - 3
	divisao := 30 / 2
	multiplicacao := 2 * 3
	restoDaDiviscao := 30 % 2

	fmt.Println(soma, subtrcao, divisao, multiplicacao, restoDaDiviscao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)
	//FIM DOS OPERADORES ARITMETICOS

	//ATRIBUICAO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	//FIM DA ATRIBUICAO

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	//FIM DOS OPERADORES RELACIONAIS

	//OPERADORES LOGICOS
	fmt.Println("------------------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // all true return true else false
	fmt.Println(verdadeiro || falso) // one true return true else, else false
	fmt.Println(!verdadeiro)         // invert the value
	fmt.Println(!falso)
	//FIM DOS OPERADORES LOGICOS

	//OPERADORES UNARIOS
	fmt.Println("------------------------------------")
	numero := 10
	numero++
	numero += 10 // numero = numero +10
	fmt.Println(numero)
	numero--
	numero -= 10 // numero = numero -10
	fmt.Println(numero)
	//FIM DOS OPERADORES UNARIOS

	var texto string

	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	//texto := numero > 5 ? "Maior que 5" : "Menor que 5" in go there is no ternary operator
	fmt.Println(texto)

}

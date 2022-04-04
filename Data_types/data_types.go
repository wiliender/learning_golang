package main

import (
	"errors"
	"fmt"
)

func main() {
	//int8, int16, int32, int64
	var numero int16 = 100
	var numero2 int64 = -1000000000000000
	var numero3 int = 100
	numero4 := 10000000 // inferencia de tipo
	fmt.Println(numero)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)

	//uint //unsigned int
	var numero5 uint32 = 100000
	fmt.Println(numero5)

	//alias
	//INT32 = RUNE
	var numero6 rune = 12456
	fmt.Println(numero6)

	//BYTE = UINT8
	var numero7 byte = 123
	fmt.Println(numero7)
	//FIM NUMEROS INTEIROS

	//NUMEROS REAIS

	//float32, float64
	var numeroReal1 float32 = 12.345
	fmt.Println(numeroReal1)

	var numeroReal2 float32 = 120000000000000000.12345
	fmt.Println(numeroReal2)

	numeroReal3 := 12.345
	fmt.Println(numeroReal3)

	//FIM NUMEROS REAIS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// go doesn't have char
	char := 'B'
	fmt.Println(char)

	//FIM STRINGS

	var texto string
	fmt.Println(texto)

	var texto2 int
	fmt.Println(texto2)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro1 error
	fmt.Println(erro1)

	var erro error = errors.New("Erro interno.")
	fmt.Println(erro)
}

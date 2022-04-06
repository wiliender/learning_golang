package main

import "fmt"

func main() {
	fmt.Println("Control Structures")

	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//when you create a variable using if init it is limited to the scope of if
	if numero2 := numero; numero2 > 0 {
		fmt.Println("Maior que 0")
	} else if numero < -10 {
		fmt.Println("Menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}

	//fmt.Println(numero2) this variable is not available outside the scope of if
}

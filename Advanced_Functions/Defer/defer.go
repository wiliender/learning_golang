package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a Funcao 1")
}

func funcao2() {
	fmt.Println("Executando a Funcao 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. Resultado será retornado")
	fmt.Println("Executando a Funcao para verificar se o  aluno esta aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		//fmt.Println("Media calculada. Resultado será retornado")
		return true
	}
	//fmt.Println("Media calculada. Resultado será retornado")
	return false
}

func main() {
	/* defer funcao1()
	// DEFER = ADIAR
	funcao2() */
	fmt.Println(alunoEstaAprovado(5, 5))
}

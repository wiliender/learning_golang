package main

import "fmt"

func recuperarExecucao() {
	//fmt.Println("Tentativa de recuperar a execução")
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A MEDIA É EXATAMENTE 6!") //INTERROMPE O PROGRAMA E PARA TUDO OQ TA EXECUTANDO
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução do programa")
}

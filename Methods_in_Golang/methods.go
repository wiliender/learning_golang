package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário: %s\n", u.nome)
}

func (u usuario) maiorDeIdade() {
	if u.idade >= 18 {
		fmt.Printf("O usuário %s é maior de idade\n", u.nome)
	} else {
		fmt.Printf("O usuário %s é menor de idade\n", u.nome)
	}
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Wiliender", 24}
	usuario1.salvar()
	usuario1.maiorDeIdade()

	usuario2 := usuario{"Daniel", 30}
	usuario2.salvar()
	usuario2.maiorDeIdade()

	usuario3 := usuario{"João", 15}
	usuario3.salvar()
	usuario3.maiorDeIdade()

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}

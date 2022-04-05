package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	fmt.Println(u)
	u.nome = "Wiliender"
	u.idade = 30
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua exemplo", 123}

	usuario2 := usuario{"Wiliender", 24, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Wiliender"}
	fmt.Println(usuario3)

	usuario4 := usuario{idade: 24}
	fmt.Println(usuario4)
}

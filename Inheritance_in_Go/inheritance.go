package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

//Go does not support inheritance. However it does support composition
func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Wiliender", "Silva", 24, 180}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia da Computação", "FHO"}
	fmt.Println(e1)
}

package main

import (
	"fmt"
	"puppy/Testes_Automatizados/introducao/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
}

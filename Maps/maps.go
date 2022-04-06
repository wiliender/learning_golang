package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Wiliender",
		"sobreNome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiroNome": "Wiliender",
			"segundoNome":  "Silva",
		},
		"curso": {
			"nome":       "Aprenda Golang do Zero!",
			"plataforma": "Udemy",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["apelido"] = map[string]string{
		"nome": "Wil",
	}
	fmt.Println(usuario2)
}

package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page!"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregando página de usuários..."))
}

func main() {
	//HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

	//CLIENTE - SERVIDOR

	// Request - Response

	// Rotas
	// URI - Identificador do Recurso
	// Método HTTP - GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)
	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

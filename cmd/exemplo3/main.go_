package main

import (
	"fmt"
	"net/http"

	"github.com/ferfabricio/golang-workshop/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /produtos", handlers.NewListaProdutosHandler())

	fmt.Println("Servidor rodando na porta 8123")
	err := http.ListenAndServe(":8123", mux)
	if err != nil {
		panic(err)
	}
}

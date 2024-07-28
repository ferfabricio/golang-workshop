package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/ferfabricio/golang-workshop/internal/handlers"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /produtos", handlers.NewListaProdutosHandler(logger))
	mux.HandleFunc("POST /ordens", handlers.NewCriarOrdemHandler(logger))

	fmt.Println("Servidor rodando na porta 8123")
	err := http.ListenAndServe(":8123", mux)
	if err != nil {
		panic(err)
	}
}

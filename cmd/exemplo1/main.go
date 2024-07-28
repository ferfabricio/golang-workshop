package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. criar um servidor http
	// 2. criar uma rota de hello world

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`Hello world`))
		if err != nil {
			panic(err)
		}
	}))

	fmt.Print("Servidor iniciado na porta 8123")
	err := http.ListenAndServe(":8123", mux)
	if err != nil {
		panic(err)
	}
}

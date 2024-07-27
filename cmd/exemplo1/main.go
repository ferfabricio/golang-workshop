package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá, mundo!"))
	})

	err := http.ListenAndServe(":8123", mux)
	if err != nil {
		panic(err)
	}
}

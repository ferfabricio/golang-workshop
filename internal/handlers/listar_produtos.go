package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ferfabricio/golang-workshop/internal/produtos"
)

func NewListaProdutosHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		produtos := produtos.CarregarLista()

		data, err := json.Marshal(produtos)
		if err != nil {
			http.Error(w, "Erro ao serializar produtos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(data)
	}
}

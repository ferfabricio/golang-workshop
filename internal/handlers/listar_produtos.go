package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ferfabricio/golang-workshop/internal/produtos"
)

func NewListaProdutosHandler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		produtos := produtos.CarregarLista()

		data, err := json.Marshal(produtos)
		if err != nil {
			http.Error(w, "Erro ao serializar produtos", http.StatusInternalServerError)
			return
		}

		logger.Info("produtos na resposta", "quantidade", len(produtos), "agent", r.Header.Get("User-Agent"))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(data)
	}
}

package handlers

import (
	"encoding/json"
	"log/slog"

	"io"
	"net/http"

	"github.com/ferfabricio/golang-workshop/internal/types"
)

func NewCriarOrdemHandler(logger *slog.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Erro ao ler corpo da requisição", http.StatusInternalServerError)
			return
		}

		var ordem types.Ordem
		err = json.Unmarshal(body, &ordem)
		if err != nil {
			logger.Error("erro ao deserializar body", "error", err)
			http.Error(w, "Erro ao deserializar ordem", http.StatusInternalServerError)
			return
		}

		// processar o ordem em algum lugar
		logger.Info("orderm recebida", "cliente", ordem.Cliente, "produtos", ordem.Produtos)

		response := `{"id": 1, "status": "criada"}`

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(response))
	}
}

package main

import (
	"encoding/json"
	"fmt"
)

type Pedido struct {
	Cliente  string    `json:"cliente"`
	Produtos []Produto `json:"produtos"`
}

type Produto struct {
	Nome       string   `json:"primeiro_nome"`
	Quantidade int      `json:"quantidade"`
	Valor      float64  `json:"-"`
	Cores      []string `json:"cores"`
}

// Objetivos:
// 1. Modificar os struct para mapear os campos json
// 2. serializar o objeto json
// 3. deserializar o objeto json
func main() {
	pedido := Pedido{
		Cliente: "Fulano",
		Produtos: []Produto{
			{
				Nome:       "Camiseta",
				Quantidade: 2,
				Valor:      29.90,
				Cores:      []string{"preto", "branco"},
			},
			{
				Nome:       "Adesivo",
				Quantidade: 10,
				Valor:      1.50,
			},
		},
	}
	data, err := json.MarshalIndent(pedido, "", "  ")
	// data, err := json.Marshal(pedido)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

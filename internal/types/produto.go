package types

type Produto struct {
	Nome       string   `json:"nome"`
	Quantidade int      `json:"quantidade"`
	Valor      float64  `json:"valor"`
	Cores      []string `json:"cores"`
}

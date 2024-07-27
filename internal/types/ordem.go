package types

type Ordem struct {
	Cliente  string    `json:"cliente"`
	Produtos []Produto `json:"produtos"`
}

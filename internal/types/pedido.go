package types

type Pedido struct {
	Cliente  string    `json:"cliente"`
	Produtos []Produto `json:"produtos"`
}

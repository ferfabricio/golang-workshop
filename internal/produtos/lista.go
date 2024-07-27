package produtos

import "github.com/ferfabricio/golang-workshop/internal/types"

func CarregarLista() []types.Produto {
	return []types.Produto{
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
	}
}

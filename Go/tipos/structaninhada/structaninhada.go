package main

import "fmt"

type item struct {
	produtoID int
	qtd       int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{
				produtoID: 1,
				qtd:       2,
				preco:     50.99,
			},
			{
				produtoID: 2,
				qtd:       1,
				preco:     233.95,
			},
			{
				produtoID: 3,
				qtd:       5,
				preco:     69.90,
			},
		},
	}

	fmt.Printf("Valor total do pedido: R$ %.2f\n", pedido.valorTotal())
}

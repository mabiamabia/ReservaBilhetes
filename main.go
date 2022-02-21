package main

import "fmt"

type cliente struct {
	nome  string
	email string
	carrinho
}

type item struct {
	nome       string
	valor      float32
	quantidade uint
}

type carrinho struct {
	itens []item
}

func (c *carrinho) addItem(itens ...item) {
	for _, v := range itens {
		c.itens = append(c.itens, v)
	}
}
func (c carrinho) calculaPreco() float32 {
	var sum float32 = 0.0

	for _, v := range c.itens {
		sum += v.valor * float32(v.quantidade)
	}
	return sum
}

func getItens() (item, item) {

	sabao := item{
		nome:       "Omo",
		valor:      7.50,
		quantidade: 2,
	}

	arroz := item{
		nome:       "Arroz",
		valor:      22.50,
		quantidade: 3,
	}

	return sabao, arroz
}

func main() {

	meuCliente := cliente{
		nome:  "Max",
		email: "max@globo.com",
	}

	meuCliente.carrinho.addItem(getItens())

	fmt.Println(meuCliente)
	fmt.Println("Preço total da compra do", meuCliente.nome, ":", meuCliente.carrinho.calculaPreco())

}
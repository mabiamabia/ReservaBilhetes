package main

import "fmt"

type cliente struct {
	nome  string
	email string
}

/* type bilhete struct {
	nome       string
	valor      float32
	quantidade uint
} */

func main() {
	meuCliente := cliente{
		nome:  "Pedro",
		email: "Pedro@globo.com",
	}

}

fmt.Println("teste", meuCliente.nome, meuCliente.email, "teste")


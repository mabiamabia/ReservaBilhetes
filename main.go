package main

import (
	"fmt"
)

func main() {
	var nome string
	var numberTicket float32
	var ticket float32 = 10.5
	var novaCompra float32

	for {
		novaCompra = 1

		fmt.Println("Qual é o seu nome ?")
		fmt.Scanln(&nome)
		fmt.Print("Quantos tickets você quer reservar? \n")
		fmt.Scanf("%f\n", &numberTicket)
		fmt.Println("Deseja mesmo reservar?\n", numberTicket)
		fmt.Print(nome, " você comprou ", numberTicket, " ticket(s)")
		fmt.Printf(" no valor de R$ %.2f \n", (numberTicket * ticket))
		fmt.Println("Parabéns! Sua compra foi confirmada")
		fmt.Println("Deseja realizar uma nova compra?\nDigite: 1 - Sim  2 - Não ")
		fmt.Scanln(&novaCompra)

		//if nova == s {
		/* volta ao inicio */
		//} else {
		/* Sai do sistema */
		//}

	}
}

package main

import "fmt"

func main() {
	var nome string
	var numberTicket float32
	var multiplica float32
	multiplica = numberTicket * 10.5

	fmt.Println("Qual é o seu nome ?")
	fmt.Scanln(&nome)
	fmt.Print("Quantos tickets você quer reservar? \n")
	fmt.Scanf("%f\n", &numberTicket)
	fmt.Printf("Deseja mesmo reservar %f ?\n", numberTicket)

	fmt.Println(nome, ", você reservou:", numberTicket, "ticket(s), no valor de R$ ", multiplica)
}

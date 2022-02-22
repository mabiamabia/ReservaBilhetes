package main

import "fmt"

func main() {
	var nome string
	var numberTicket float32
	fmt.Println("Qual é o seu nome ?")
	fmt.Scanln(&nome)
	fmt.Print("Quantos tickets você quer reservar? \n")
	fmt.Scanf("%f\n", &numberTicket)
	fmt.Printf(" deseja mesmo reservar %f", numberTicket)
}

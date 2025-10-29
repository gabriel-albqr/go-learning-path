package main

import "fmt"

func main() {
	numerodebytes, erros := fmt.Println("Olá, Golang!")
	fmt.Println("Número de bytes escritos:", numerodebytes)
	fmt.Println("Erros:", erros)
}

func exemplo() {
	// Declaração de variáveis
	var idade int = 25
	nome := "Gabriel"
	var altura float64 = 1.65
	var ativo bool = true
	const pi float64 = 3.14

	// Impressão dos valores
	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura)
	fmt.Println("Ativo:", ativo)
	fmt.Println("Valor de Pi:", pi)
}

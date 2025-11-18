package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Saudacao() {
	fmt.Println("Olá,", p.Nome, "- idade:", p.Idade)
}

func main() {
	pessoa := Pessoa{"Gabriel", 21}
	pessoa.Saudacao()
}

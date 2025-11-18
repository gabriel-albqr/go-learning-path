package main

import "fmt"

type Forma interface {
	Area() float64
}

type Quadrado struct {
	Lado float64
}

func (q Quadrado) Area() float64 {
	return q.Lado * q.Lado
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	const pi = 3.14
	return pi * c.Raio * c.Raio
}

func main() {
	formas := []Forma{
		Quadrado{Lado: 5},
		Circulo{Raio: 3},
	}

	for _, f := range formas {
		fmt.Println("Área:", f.Area())
	}
}

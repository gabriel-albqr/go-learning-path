package main

import "fmt"

func main() {
	x := 10
	p := &x // ponteiro para x

	fmt.Println("Valor de x:", x)
	fmt.Println("Endereço de x:", p)
	fmt.Println("Valor apontado por p:", *p)

	// Modificando o valor via ponteiro
	*p = 20
	fmt.Println("Valor modificado de x via ponteiro:", x)
}

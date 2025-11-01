package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Ímpar")
	}
}

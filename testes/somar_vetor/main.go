package main

import "fmt"

func main() {
	numeros := []int{2, 4, 6, 8, 10}
	soma := 0

	for _, v := range numeros {
		soma += v
	}

	fmt.Println("Soma total:", soma)
}

package main

import (
	"fmt"
	"time"
)

func contar(nome string, ch chan string) {
	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("%s contando: %d", nome, i)
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func main() {
	ch := make(chan string)

	go contar("Goroutine 1", ch)

	for msg := range ch {
		fmt.Println(msg)
	}

	fmt.Println("Fim do programa.")
}

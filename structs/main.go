package main

import (
	"fmt"
	"time"
)

func tarefa(nome string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(nome, i)
		time.Sleep(400 * time.Millisecond)
	}
}

func main() {
	go tarefa("Processo A")
	go tarefa("Processo B")

	time.Sleep(2 * time.Second) // espera as goroutines
}

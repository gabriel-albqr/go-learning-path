package main

import (
	"fmt"
	"time"
)

func main() {
	var segundos int

	fmt.Println("=== TEMPORIZADOR ===")
	fmt.Print("Digite o tempo (em segundos): ")
	fmt.Scan(&segundos)

	fmt.Println("\nIniciando temporizador...\n")

	for segundos >= 0 {
		fmt.Printf("\rRestam: %02d:%02d", segundos/60, segundos%60)
		time.Sleep(1 * time.Second)
		segundos--
	}

	fmt.Println("\n⏰ Tempo esgotado!")
}

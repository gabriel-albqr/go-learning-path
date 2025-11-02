package main

import (
	"fmt"
	"time"
)

func main() {
	var segundos int

	fmt.Println("=== CRONÔMETRO ===")
	fmt.Print("Digite o tempo inicial (em segundos): ")
	fmt.Scan(&segundos)

	fmt.Println("\nIniciando cronômetro... (Ctrl + C para parar)\n")

	for {
		fmt.Printf("\rTempo: %02d:%02d", segundos/60, segundos%60)
		time.Sleep(1 * time.Second)
		segundos++
	}
}

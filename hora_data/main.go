package main

import (
	"fmt"
	"time"
)

func main() {
	// obtém a data e hora atual
	now := time.Now()

	// exibe a data e hora completa
	fmt.Println("Data e hora atual:", now)

	// formatos personalizados
	fmt.Println("Formato BR:", now.Format("02/01/2006 15:04:05"))
	fmt.Println("Apenas data:", now.Format("02/01/2006"))
	fmt.Println("Apenas hora:", now.Format("15:04:05"))
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var opcao int
	for {
		fmt.Println("\n=== Conversor Binário ↔ Decimal ===")
		fmt.Println("1 - Binário para Decimal")
		fmt.Println("2 - Decimal para Binário")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			converterBinarioParaDecimal()
		case 2:
			converterDecimalParaBinario()
		case 0:
			fmt.Println("Saindo... Até logo!")
			return
		default:
			fmt.Println("Opção inválida! Tente novamente.")
		}
	}
}

func converterBinarioParaDecimal() {
	var binario string
	fmt.Print("Digite o número binário: ")
	fmt.Scan(&binario)

	decimal, err := strconv.ParseInt(binario, 2, 64)
	if err != nil {
		fmt.Println("Erro: número binário inválido!")
		return
	}
	fmt.Printf("Decimal: %d\n", decimal)
}

func converterDecimalParaBinario() {
	var decimal int64
	fmt.Print("Digite o número decimal: ")
	fmt.Scan(&decimal)

	binario := strconv.FormatInt(decimal, 2)
	fmt.Printf("Binário: %s\n", binario)
}

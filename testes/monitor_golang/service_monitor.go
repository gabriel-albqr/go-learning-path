package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	fmt.Println(" Monitor de Serviços Simples ")

	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("sc", "query", "state=", "all")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Erro ao listar serviços:", err)
			return
		}
		fmt.Println(string(output))

	case "linux":
		cmd := exec.Command("systemctl", "list-units", "--type=service", "--no-pager", "--no-legend")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Erro ao listar serviços:", err)
			return
		}
		fmt.Println(string(output))

	case "darwin":
		cmd := exec.Command("launchctl", "list")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Erro ao listar serviços:", err)
			return
		}
		fmt.Println(string(output))

	default:
		fmt.Println("Sistema operacional não suportado:", runtime.GOOS)
	}
}

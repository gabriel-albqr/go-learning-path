package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	quotes := []string{
		"“Falar é fácil. Mostre-me o código.” — Linus Torvalds",
		"“Programas devem ser escritos para que pessoas os leiam, e apenas incidentalmente para que máquinas os executem.” — Harold Abelson",
		"“A simplicidade é a alma da eficiência.” — Austin Freeman",
		"“Primeiro resolva o problema. Depois escreva o código.” — John Johnson",
		"“O código nunca mente, mas os comentários às vezes sim.” — Ron Jeffries",
		"“Corrija a causa, não o sintoma.” — Steve Maguire",
		"“Antes que o software possa ser reutilizável, ele precisa ser utilizável.” — Ralph Johnson",
		"“Faça funcionar, faça direito, depois faça rápido.” — Kent Beck",
		"“A coisa mais desastrosa que você pode aprender é a sua primeira linguagem de programação.” — Alan Kay",
		"“Qualquer tolo pode escrever código que um computador entenda. Bons programadores escrevem código que humanos entendem.” — Martin Fowler",
		"“Depurar é o dobro de difícil que escrever o código em primeiro lugar. Portanto, se você escrever o código da maneira mais inteligente possível, por definição, você não é inteligente o suficiente para depurá-lo.” — Brian Kernighan",
		"“Não é um bug — é um recurso!” — Autor desconhecido",
		"“Código limpo sempre parece que foi escrito por alguém que se importava.” — Michael Feathers",
		"“A melhor forma de prever o futuro é implementá-lo.” — David Heinemeier Hansson",
		"“Programar é o ato de transformar café em código.” — Autor desconhecido",
	}

	// somente para gerar valores aleatórios diferentes a cada execução
	rand.Seed(time.Now().UnixNano())

	// seleciona uma frase aleatória
	randomIndex := rand.Intn(len(quotes))
	fmt.Println(quotes[randomIndex])
}

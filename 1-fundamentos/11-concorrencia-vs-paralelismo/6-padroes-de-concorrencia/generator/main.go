package main

import (
	"fmt"
	"time"
)

// O padrão Generator é uma maneira de criar funções que produzem valores em um canal de forma incremental.
// Ele gera uma sequência de valores que podem ser consumidos por outras goroutines conforme necessário.
func main() {
	canal := escrever("Olá mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// O padrão Multiplexador, ou Mux, é uma forma de combinar múltiplos canais de entrada em um único canal de saída.
// Ele permite que uma goroutine possa escutar e tratar dados de vários canais ao mesmo tempo, enviando todos os valores para um único canal de saída.
func main() {
	canal := multiplexar(escrever("Olá mundo!"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return canal
}

func multiplexar(canal1, canal2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canalSaida <- mensagem
			case mensagem := <-canal2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
}

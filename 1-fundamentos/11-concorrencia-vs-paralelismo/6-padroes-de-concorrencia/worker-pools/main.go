package main

import "fmt"

// O padrão Worker-Pool é utilizado para distribuir a carga de trabalho entre várias goroutines.
// Ele cria um conjunto de goroutines que processam tarefas simultaneamente, equilibrando a carga entre elas.
// Isso é útil para realizar operações intensivas, como processamento paralelo de tarefas em lotes.
func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	defer close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

// <-chan (somente recebe dados)
// chan<- (somente envia dados)
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

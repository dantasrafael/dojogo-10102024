package main

import (
	"fmt"
	"time"
)

// Um canal em Go é uma estrutura que permite a comunicação segura entre goroutines.
// Eles são utilizados para enviar e receber dados entre goroutines,
// sincronizando sua execução e permitindo o compartilhamento de informações sem a necessidade de mecanismos mais complexos como locks (travas).
func main() {
	canal := make(chan string)

	go escrever("Olá Mundo!", canal)
	fmt.Println("Depois de escrever")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}

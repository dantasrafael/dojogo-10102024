package main

import (
	"fmt"
	"time"
)

// O select em Go serve para realizar operações de comunicação com múltiplos canais simultaneamente,
// permitindo que uma goroutine espere em várias operações de canais ao mesmo tempo.
// O select escolhe a primeira operação de canal que estiver pronta e a executa, seja para enviar ou receber dados.
// Se várias operações estiverem prontas, o select escolhe uma delas de forma aleatória.
func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "canal 2"
		}
	}()

	for {
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}

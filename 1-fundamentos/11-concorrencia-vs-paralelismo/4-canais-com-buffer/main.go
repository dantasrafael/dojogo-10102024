package main

import "fmt"

// Canais com buffer em Go permitem o envio de um número limitado de valores sem bloquear imediatamente a goroutine que está enviando.
// Diferente dos canais sem buffer (onde cada envio bloqueia até que o valor seja recebido),
// um canal bufferizado permite que várias mensagens sejam enviadas antes que o receptor precise consumi-las, até o limite do buffer.
func main() {
	canal := make(chan string, 2)
	fmt.Println(cap(canal))

	fmt.Println(len(canal))

	canal <- "Olá mundo!"
	canal <- "Programando em Go!"

	fmt.Println(len(canal))

	mensagem1 := <-canal
	mensagem2 := <-canal

	fmt.Println(len(canal))

	fmt.Println(mensagem1)
	fmt.Println(mensagem2)
}

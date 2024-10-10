package main

import (
	"desafio-dantas/cmd"
	"fmt"
	"math/rand"
	"time"
)

// Função auxiliar para gerar um pedido com itens aleatórios.
func gerarPedido(id int, isVIP bool) cmd.Pedido {
	numItems := rand.Intn(3) + 1
	items := make([]cmd.Item, numItems)
	for i := range items {
		items[i] = cmd.Item{Nome: fmt.Sprintf("Item-%d", rand.Intn(100))}
	}
	return cmd.Pedido{
		ID:     id,
		Items:  items,
		Pronto: false,
		IsVIP:  isVIP,
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	restaurante := cmd.NovoRestaurante(10)

	// Simula recebimento de pedidos
	for i := 1; i <= 5; i++ {
		isVIP := rand.Float32() < 0.2 // 20% de chance de ser VIP
		pedido := gerarPedido(i, isVIP)
		restaurante.ReceberPedido(pedido)
	}

	// Processamento dos pedidos em paralelo
	go restaurante.ProcessarPedidos()

	// Aguarda por um tempo antes de encerrar o restaurante
	time.Sleep(10 * time.Second)
	restaurante.FecharRestaurante()
}

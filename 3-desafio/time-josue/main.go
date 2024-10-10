package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

type Item struct {
	nome string
}

type Pedido struct {
	nome   string
	tempo  int
	status string
	Itens  []Item
}

func main() {
	pedidos := listaPedidos(pedidos(), pedidosVips())

	for pedido := range pedidos {
		processaPedido(pedido)
	}

}

func pedidos() <-chan Pedido {
	c := make(chan Pedido)
	go func() {
		i := 1
		for {
			c <- Pedido{
				nome:   fmt.Sprintf("Pedido %d", i),
				tempo:  1000,
				status: "pendente",
				Itens: []Item{
					{nome: "Item 1"},
					{nome: "Item 2"},
				},
			}
			time.Sleep(time.Second * 2)
			i++

		}
	}()
	return c
}

func pedidosVips() <-chan Pedido {
	c := make(chan Pedido)
	go func() {
		i := 1
		for {
			c <- Pedido{
				nome:   fmt.Sprintf("Pedido VIP %d", i),
				tempo:  1000,
				status: "pendente",
				Itens: []Item{
					{nome: "Item 1"},
					{nome: "Item 2"},
				},
			}
			time.Sleep(time.Second * 3)
			i++
		}
	}()
	return c
}

func processaPedido(pedido Pedido) {
	fmt.Printf("\n\nProcessando Pedido: %v\n", pedido.nome)
	for _, item := range pedido.Itens {
		waitGroup.Add(1)
		go processaItem(item)
	}

	time.Sleep(time.Millisecond * time.Duration(pedido.tempo))

	waitGroup.Wait()
	pedido.status = "finalizado"
}

func processaItem(item Item) {
	defer waitGroup.Done()
	fmt.Printf("Preparando item: %v\n", item.nome)
}

func listaPedidos(canalComum, canalVip <-chan Pedido) <-chan Pedido {
	canalSaida := make(chan Pedido)

	go func() {
		for {
			select {
			case pedido := <-canalVip:
				canalSaida <- pedido
			case pedido := <-canalComum:
				canalSaida <- pedido
			}
		}
	}()

	return canalSaida
}

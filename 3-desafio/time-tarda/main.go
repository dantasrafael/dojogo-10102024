package main

import (
	"fmt"
	"github.com/govalues/decimal"
	"sync"
	"time"
)

type Item struct {
	Nome       string
	Quantidade int32
	Valor      decimal.Decimal
}

type Pedido struct {
	Id          int
	Itens       []Item
	DtCriacao   time.Time
	Prioritario bool
}

func MontarPedido(itens []Item, prioritario bool, id int) *Pedido {
	return &Pedido{
		Id:          id,
		Itens:       itens,
		DtCriacao:   time.Now(),
		Prioritario: prioritario,
	}
}

func mockPedidos() []Pedido {
	pedido1 := MontarPedido([]Item{
		{Nome: "Batata", Quantidade: 2, Valor: decimal.MustParse("15.11")},
		{Nome: "Frango", Quantidade: 3, Valor: decimal.MustParse("122.11")},
		{Nome: "Cebola outback", Quantidade: 1, Valor: decimal.MustParse("50.11")},
	}, false, 1)

	pedido2 := MontarPedido([]Item{{Nome: "Batata", Quantidade: 1, Valor: decimal.MustParse("15.11")}}, false, 2)

	pedido3 := MontarPedido([]Item{{Nome: "Frango", Quantidade: 1, Valor: decimal.MustParse("22.11")}}, false, 3)

	pedido4 := MontarPedido([]Item{{Nome: "Costela", Quantidade: 1, Valor: decimal.MustParse("33.11")}}, true, 4)

	pedido5 := MontarPedido([]Item{{Nome: "Ancho", Quantidade: 1, Valor: decimal.MustParse("44.11")}}, true, 5)

	return []Pedido{*pedido1, *pedido2, *pedido3, *pedido4, *pedido5}
}

func processarPedido(pedido *Pedido) {
	fmt.Printf("Processando pedido Id: %d \n", pedido.Id)
	var wg sync.WaitGroup

	var total decimal.Decimal
	for _, item := range pedido.Itens {
		wg.Add(1)
		go processarItem(item, &wg, &total, pedido.Id)
	}

	wg.Wait()

	fmt.Printf("Total do pedido %d Valor: U$ %s \n", pedido.Id, total.String())
	fmt.Printf("Pedido processado: %d, qntd de itens do pedido: %d \n\n", pedido.Id, len(pedido.Itens))
}

func processarItem(item Item, wg *sync.WaitGroup, total *decimal.Decimal, pedido int) {
	defer wg.Done()

	fmt.Printf("Pedido: %d, Processando item: %s, qntd: %d, valor: U$ %s \n", pedido, item.Nome, item.Quantidade, item.Valor.String())
	time.Sleep(1000 * time.Millisecond)
	*total, _ = total.Add(item.Valor)
}

func main() {
	timeout := time.After(5 * time.Second)
	done := make(chan bool)

	pedidos := mockPedidos()
	for _, pedido := range pedidos {

		go func() {
			processarPedido(&pedido)
			done <- true

			select {
			case <-done:
				fmt.Printf("Pedido #%d concluído dentro do tempo!\n", pedido.Id)
			case <-timeout:
				fmt.Printf("Pedido #%d não foi concluído a tempo!\n", pedido.Id)
			}
		}()

	}

	time.Sleep(10 * time.Second)
}

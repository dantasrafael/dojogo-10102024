package cmd

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Restaurante define o fluxo de recebimento e processamento dos pedidos.
type Restaurante struct {
	pedidosCh chan Pedido
	wgPedidos sync.WaitGroup
}

// NovoRestaurante cria uma nova instância de Restaurante.
func NovoRestaurante(buffer int) *Restaurante {
	return &Restaurante{
		pedidosCh: make(chan Pedido, buffer),
	}
}

// RecebePedido adiciona um novo pedido à fila para ser processado.
func (r *Restaurante) ReceberPedido(pedido Pedido) {
	r.wgPedidos.Add(1)
	go func() {
		r.pedidosCh <- pedido
	}()
}

// ProcessaPedidos processa todos os pedidos simultaneamente usando goroutines.
func (r *Restaurante) ProcessarPedidos() {
	for pedido := range r.pedidosCh {
		go r.processarPedido(pedido)
	}
}

// FechaRestaurante fecha o canal de pedidos e aguarda o término do processamento.
func (r *Restaurante) FecharRestaurante() {
	close(r.pedidosCh)
	r.wgPedidos.Wait()
	fmt.Println("Todos os pedidos foram processados!")
}

// processaPedido processa cada item de um pedido.
func (r *Restaurante) processarPedido(pedido Pedido) {
	defer r.wgPedidos.Done()

	fmt.Printf("Processando pedido #%d com %d itens...\n", pedido.ID, len(pedido.Items))

	var wgItems sync.WaitGroup
	for _, item := range pedido.Items {
		wgItems.Add(1)
		go r.prepararItem(pedido.ID, item, &wgItems)
	}

	wgItems.Wait()
	fmt.Printf("Pedido #%d pronto!\n", pedido.ID)
}

// prepararItem simula o preparo de um item do pedido.
func (r *Restaurante) prepararItem(pedidoID int, item Item, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simula o tempo de preparo do item
	tempoPreparo := rand.Intn(3) + 1
	time.Sleep(time.Duration(tempoPreparo) * time.Second)

	fmt.Printf("Item '%s' do pedido #%d está pronto! (Tempo: %ds)\n", item.Nome, pedidoID, tempoPreparo)
}

package cmd

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const pedidoTimeout = 5 * time.Second

// Restaurante define o fluxo de recebimento e processamento dos pedidos.
type Restaurante struct {
	pedidosVIPCh  chan Pedido
	pedidosNormCh chan Pedido
	wgPedidos     sync.WaitGroup
}

// NovoRestaurante cria uma nova instância de Restaurante.
func NovoRestaurante(buffer int) *Restaurante {
	return &Restaurante{
		pedidosVIPCh:  make(chan Pedido, buffer),
		pedidosNormCh: make(chan Pedido, buffer),
	}
}

// ReceberPedido adiciona um novo pedido à fila para ser processado.
func (r *Restaurante) ReceberPedido(pedido Pedido) {
	r.wgPedidos.Add(1)

	go func() {
		if pedido.IsVIP {
			fmt.Printf("Recebido pedido VIP #%d\n", pedido.ID)
			r.pedidosVIPCh <- pedido
		} else {
			fmt.Printf("Recebido pedido normal #%d\n", pedido.ID)
			r.pedidosNormCh <- pedido
		}
	}()
}

// ProcessarPedidos processa todos os pedidos simultaneamente usando goroutines.
func (r *Restaurante) ProcessarPedidos() {
	go r.processarPedidosVIP()
	go r.processarPedidosNormais()
}

// FecharRestaurante fecha o canal de pedidos e aguarda o término do processamento.
func (r *Restaurante) FecharRestaurante() {
	close(r.pedidosVIPCh)
	close(r.pedidosNormCh)
	r.wgPedidos.Wait()

	fmt.Println("Todos os pedidos foram processados!")
}

// processarPedidosVIP processa os pedidos VIP.
func (r *Restaurante) processarPedidosVIP() {
	for pedido := range r.pedidosVIPCh {
		go r.processarPedidoComTimeout(pedido)
	}
}

// processarPedidosNormais processa os pedidos normais.
func (r *Restaurante) processarPedidosNormais() {
	for pedido := range r.pedidosNormCh {
		go r.processarPedidoComTimeout(pedido)
	}
}

// processarPedidoComTimeout processa cada pedido e cancela se ultrapassar o timeout.
func (r *Restaurante) processarPedidoComTimeout(pedido Pedido) {
	defer r.wgPedidos.Done()

	done := make(chan bool)
	go func() {
		r.processarPedido(pedido)
		pedido.Pronto = true
		done <- true
	}()

	select {
	case <-done:
		fmt.Printf("Pedido #%d processado com sucesso.\n", pedido.ID)
	case <-time.After(pedidoTimeout):
		pedido.Expirado = true
		fmt.Printf("Pedido #%d expirado! Tempo de processamento excedido.\n", pedido.ID)
	}
}

// processaPedido processa cada item de um pedido.
func (r *Restaurante) processarPedido(pedido Pedido) {
	fmt.Printf("Processando pedido #%d com %d itens...\n", pedido.ID, len(pedido.Items))

	var wgItems sync.WaitGroup
	for _, item := range pedido.Items {
		wgItems.Add(1)
		go r.prepararItem(pedido.ID, item, &wgItems)
	}

	wgItems.Wait()
}

// prepararItem simula o preparo de um item do pedido.
func (r *Restaurante) prepararItem(pedidoID int, item Item, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simula o tempo de preparo do item
	tempoPreparo := rand.Intn(3) + 1
	time.Sleep(time.Duration(tempoPreparo) * time.Second)

	fmt.Printf("Item '%s' do pedido #%d está pronto! (Tempo: %ds)\n", item.Nome, pedidoID, tempoPreparo)
}

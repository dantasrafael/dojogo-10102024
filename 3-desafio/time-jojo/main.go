package main

import (
	"sync"
	"time"

	"github.com/fatih/color"
)

type Status int

const (
	Pronto Status = iota
	EmPreparo
	EmEspera
)

var green = color.New(color.FgGreen).PrintfFunc()
var yellow = color.New(color.FgYellow).PrintfFunc()

type Cozinha struct {
	Capacidade int
	EmExecucao int

}

func (c *Cozinha) prepararPedido(pedido *Pedido, wg *sync.WaitGroup) {
	yellow("Preparando pedido %d\n", pedido.Id)
	pedido.Status = EmPreparo

	for _, item := range pedido.Items {
		go c.prepararItem(item, &pedido.TotalDeItems)
	}

	pedido.TotalDeItems.Wait()

	green("Pedido %d pronto\n", pedido.Id)
	pedido.Status = Pronto
	wg.Done()
}

func (c *Cozinha) prepararItem(item *Item, wg *sync.WaitGroup) {
	yellow("Preparando Item %s\n", item.Nome)
	item.Status = EmPreparo

	time.Sleep(time.Second * time.Duration(item.TempoDePreparo))

	green("Item %s pronto\n", item.Nome)
	item.Status = Pronto
	wg.Done()
}

type Restaurante struct {
	Cozinha         *Cozinha
	NumeroDePedidos sync.WaitGroup
	garcon          chan *Pedido
}

func (r *Restaurante) AdicionarPedido(pedido *Pedido) {
	r.NumeroDePedidos.Add(1)
	r.garcon <- pedido

}

func (r *Restaurante) FecharRestaurante() {
	r.NumeroDePedidos.Wait()
	close(r.garcon)
}

func (r *Restaurante) AbriCozinha() {
	r.Cozinha.Capacidade++
}

func (r *Restaurante) BotarGarconParaTrabalhar() {
	for pedido := range r.garcon {
		go r.Cozinha.prepararPedido(pedido, &r.NumeroDePedidos)
	}
}

type Pedido struct {
	Id           int
	Items        []*Item
	Vip          bool
	Status       Status
	TotalDeItems sync.WaitGroup
}

func (p *Pedido) AdicionarItem(item *Item) {
	p.TotalDeItems.Add(1)
	p.Items = append(p.Items, item)
}

type Item struct {
	Nome           string
	TempoDePreparo int
	Status         Status
}

func NewBuchada() *Item {
	return &Item{
		Nome:           "Buchada",
		TempoDePreparo: 2,
		Status:         EmEspera,
	}
}
func NewTapiocaDeCoco() *Item {
	return &Item{
		Nome:           "Tapioca de Coco",
		TempoDePreparo: 5,
		Status:         EmEspera,
	}
}
func NewCaldoDeCana() *Item {
	return &Item{
		Nome:           "Caldo de Cana",
		TempoDePreparo: 2,
		Status:         EmEspera,
	}
}
func NewAguaTonica() *Item {
	return &Item{
		Nome:           "Água Tônica",
		TempoDePreparo: 1,
		Status:         EmEspera,
	}
}
func NewCombucha() *Item {
	return &Item{
		Nome:           "Combucha",
		TempoDePreparo: 10,
		Status:         EmEspera,
	}
}

func main() {
	BudegaDoBileu := Restaurante{
		Cozinha: &Cozinha{
			Capacidade: 1,
		},
		garcon: make(chan *Pedido),
	}
	go BudegaDoBileu.BotarGarconParaTrabalhar()

	pedido1 := Pedido{}
	pedido1.AdicionarItem(NewTapiocaDeCoco())
	pedido1.AdicionarItem(NewCaldoDeCana())
	BudegaDoBileu.AdicionarPedido(&pedido1)

	pedido2 := Pedido{
		Id: 1,
	}
	pedido2.AdicionarItem(NewBuchada())
	pedido2.AdicionarItem(NewAguaTonica())
	BudegaDoBileu.AdicionarPedido(&pedido2)

	pedido3 := Pedido{
		Id:  2,
		Vip: true,
	}
	pedido3.AdicionarItem(NewCombucha())
	BudegaDoBileu.AdicionarPedido(&pedido3)


	pedido4 := Pedido{
		Id:  3,
		Vip: true,
	}
	pedido4.AdicionarItem(NewAguaTonica())
	BudegaDoBileu.AdicionarPedido(&pedido4)

	BudegaDoBileu.FecharRestaurante()
}

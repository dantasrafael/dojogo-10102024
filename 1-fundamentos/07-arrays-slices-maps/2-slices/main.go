package main

import (
	"fmt"

	ps "github.com/inancgumus/prettyslice"
)

// Um slice é uma estrutura mais flexível que permite trabalhar com sequências de elementos de forma dinâmica.
// Embora seja baseado em arrays, um slice pode ter seu tamanho alterado dinamicamente conforme novos elementos são adicionados ou removidos.
//
// Características de um Slice:
//   - Tamanho variável: O tamanho de um slice pode mudar dinamicamente, o que o torna mais flexível que arrays.
//   - Referência a um array: Um slice não armazena diretamente os dados, mas sim uma referência a uma porção de um array subjacente.
//   - Mais comum em Go: Slices são amplamente usados em Go porque são mais práticos e dinâmicos do que arrays.
//   - Capacidade (capacity): Além do tamanho (length), slices também possuem uma capacidade (capacity), que define quantos elementos podem ser armazenados antes que seja necessário realocar um novo array subjacente.
func main() {
	fmt.Println("SLICES")

	ps.PrintBacking = true
	ps.PrintElementAddr = true
	ps.MaxPerLine = 10

	slice1 := []int{35, 15, 25}
	slice2 := slice1[:3]
	ps.Show("slice1", slice1)
	ps.Show("slice2", slice2)
	slice1[2] = 5
	ps.Show("slice1", slice1)
	ps.Show("slice2", slice2)
}

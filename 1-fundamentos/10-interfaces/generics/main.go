package main

import (
	"fmt"
)

// Define uma interface que restringe os tipos aceitos pela função Max
// Os tipos devem ser "ordenáveis", como int, float64, string, etc.
// O operador ~ permite que tipos definidos a partir de int, float64 ou string também sejam incluídos.
// Isso assegura que a função só aceite tipos que suportam as operações de comparação >.
type Ordenavel interface {
	~int | ~float64 | ~string
}

// Função genérica que encontra o valor máximo em um slice
func Max[T Ordenavel](slice []T) T {
	if len(slice) == 0 {
		panic("slice vazio")
	}

	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	// Exemplo com inteiros
	ints := []int{10, 20, 5, 40, 30}
	fmt.Println("Máximo de ints:", Max(ints)) // Saída: 40

	// Exemplo com floats
	floats := []float64{10.5, 2.3, 45.8, 3.9}
	fmt.Println("Máximo de floats:", Max(floats)) // Saída: 45.8

	// Exemplo com strings
	strings := []string{"apple", "banana", "pear", "grape"}
	fmt.Println("Máximo de strings:", Max(strings)) // Saída: pear (lexicograficamente maior)
}

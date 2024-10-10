package main

import (
	"fmt"
)

// Um array é uma estrutura de dados de tamanho fixo, onde todos os elementos têm o mesmo tipo e o tamanho é parte de sua definição.
// Isso significa que o tamanho do array é imutável após sua criação.
//
// Características de um Array:
//   - Tamanho fixo: O tamanho do array é definido no momento da sua criação e não pode ser alterado.
//   - Parte do tipo: O tamanho do array é parte do seu tipo, então [5]int e [10]int são tipos diferentes.
//   - Eficiência: Como o tamanho é conhecido, os arrays são muito eficientes em termos de alocação de memória.
func main() {
	fmt.Println("ARRAYS")

	var array1 [5]int
	array1[0] = 1
	fmt.Printf("Tipo da array1: %T\n", array1)
	fmt.Println("array1", array1)

	array2 := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("Tipo da array2: %T\n", array2)
	fmt.Println("array2", array2)

	array3 := [...]int{100, 200, 300, 400, 500}
	fmt.Printf("Tipo da array3: %T\n", array3)
	fmt.Println("array3", array3)
}

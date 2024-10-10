package main

import "fmt"

func function1() {
	fmt.Println("Function 1")
}

func function2() {
	fmt.Println("Function 2")
}

func isApprovedStudent(n1, n2 float32) bool {
	defer fmt.Printf("### Returning average = ")
	fmt.Println("### Calculating approved student")

	average := (n1 + n2) / 2
	if average >= 6 {
		return true
	}
	return false
}

// a palavra-chave 'defer' é usada para adiar a execução de uma função até que a função que a contém retorne.
// Neste caso adiamos a execução da função 'function1' até que a função 'main' esteja prestes a retornar.
// Em seguida, chama a função 'function2' e imprime o resultado da função 'isApprovedStudent' com os parâmetros 7 e 8.
func main() {
	defer function1()
	function2()

	fmt.Println(isApprovedStudent(7, 8))
}

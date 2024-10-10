package main

import "fmt"

func closure() func() {
	text := "Value inside closure function"

	return func() {
		fmt.Println(text)
	}
}

func closureNeverCalled() func() {
	fmt.Println("execution inside closure called")

	return func() {
		fmt.Println("closure never called")
	}
}

// Funções anônimas podem capturar e reter o estado de variáveis externas ao seu escopo, formando closures.
// Isso permite que você mantenha o estado entre chamadas.
func main() {
	text := "Value from main"
	fmt.Println(text)

	_ = closureNeverCalled()
	returnedFunction := closure()
	returnedFunction()
}

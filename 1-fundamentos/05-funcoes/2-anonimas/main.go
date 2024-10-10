package main

import "fmt"

// uma função anônima é uma função sem nome que pode ser declarada e usada inline
// muito útil para funções simples que não precisam ser reutilizadas.
//
// a função anônima pode ser:
//
//   - atribuída a uma variável ou chamada diretamente
//   - passada como argumento para outra função
//   - retornada por outra função
//   - usada como um tipo de dado
func main() {
	func(text string) {
		fmt.Println(text)
	}("Hello World!")

	func() {
		fmt.Println("anonymous function without parameter...")
	}()

	result := func(text string) string {
		return fmt.Sprintf("Param => %s\n", text)
	}

	fmt.Println(result("param value"))
}

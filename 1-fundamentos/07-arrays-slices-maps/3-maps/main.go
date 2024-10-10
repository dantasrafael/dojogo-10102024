package main

import "fmt"

// Map é uma estrutura de dados que associa chaves a valores.
// É uma coleção desordenada onde cada valor é acessado por meio de uma chave exclusiva, o que facilita a busca rápida por valores.
// Os mapas em Go são similares a dicionários em outras linguagens como Python ou hash tables em linguagens como Java.
//
// Características dos Maps:
//   - Chave-valor: Um map associa um valor a uma chave. As chaves devem ser de um tipo que possa ser comparado com == (como string, int, bool, etc.).
//   - Desordenado: A ordem dos pares chave-valor em um map não é garantida.
//   - Acesso eficiente: A busca por um valor, dada uma chave, é feita em tempo constante (O(1)), o que torna os mapas eficientes para esse tipo de operação.
//   - Mutável: É possível adicionar, remover ou atualizar valores em um map.
func main() {
	usuario1 := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario1)
	fmt.Println(usuario1["nome"])
	fmt.Println(usuario1["sobrenome"])

	delete(usuario1, "sobrenome")
	fmt.Println(usuario1)
	fmt.Println(usuario1["nome"])
	fmt.Println(usuario1["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nomes": {
			"primeiro": "João",
			"último":   "Pedro",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nomes"])
	fmt.Println(usuario2["nomes"]["primeiro"])

	usuario2["cursos"] = map[string]string{
		"nome":   "Medicina",
		"campus": "Centro",
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nomes"])
	fmt.Println(usuario2["nomes"]["primeiro"])
}

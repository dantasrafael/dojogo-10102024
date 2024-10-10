package main

import (
	"fmt"
	"unsafe"
)

type Pessoa struct {
	Nome  string // 16 bytes
	Idade uint8  // 1 byte
	Nick  string // 16 bytes
} // 33 bytes??? | 8 * 5 = 40 bytes

// Em Go, a escolha entre usar int e uint (inteiro sem sinal) para desempenho depende do caso de uso,
// mas as diferenças de desempenho são geralmente pequenas,
// a menos que você esteja lidando com processamento de dados em grande escala ou tarefas muito sensíveis à performance.
//
// Aqui estão algumas considerações:
//   - Faixa de valores e Segurança:
//     int: Representa inteiros com sinal e pode armazenar valores positivos e negativos. Em sistemas de 64 bits, o int é geralmente de 64 bits, e em sistemas de 32 bits, é de 32 bits.
//     uint: Representa inteiros sem sinal, armazenando apenas números não negativos. Isso oferece uma faixa um pouco maior de valores positivos, mas não permite valores negativos.
//   - Considerações de Desempenho:
//     Tamanho e Uso de Memória: Tanto o int quanto o uint ocupam a mesma quantidade de memória (por exemplo, 64 bits em sistemas de 64 bits). A diferença está na capacidade de armazenar números: uint pode armazenar números positivos maiores, enquanto int pode armazenar números positivos e negativos.
//     Operações Aritméticas: Não há uma diferença significativa no desempenho de operações aritméticas como adição, subtração ou multiplicação entre int e uint. CPUs modernas lidam bem com ambos os tipos.
//     Overflow e Underflow: Usar uint requer mais cuidado para evitar o underflow (por exemplo, ao subtrair um número maior de um menor). Isso pode adicionar complexidade ao código em alguns casos.
//     Conversões: Se o seu programa exigir muitas conversões entre int e uint, o overhead adicional dessas conversões pode impactar o desempenho de forma leve.
//   - 3. Casos de Uso:
//     Use int quando você precisar lidar com valores negativos ou não tiver certeza de que o valor sempre será positivo.
//     Use uint quando você souber que só vai trabalhar com valores positivos e precisa de um alcance maior de números positivos.
//     Em termos de desempenho, o uso de int ou uint não costuma gerar diferenças perceptíveis em operações aritméticas comuns, mas a escolha certa pode ajudar a evitar erros e otimizar a legibilidade e segurança do código.
func main() {
	var bool1 bool
	var numInt8 int8
	var numInt16 int16
	var numInt32 int32
	var numInt64 int64
	var numUint8 uint8
	var numUint16 uint16
	var numUint32 uint32
	var numUint64 uint64
	var realFloat32 float32
	var texto string
	var estrutura1 Pessoa

	arrayBool := [2]bool{true, false}
	arrayInt16 := [2]int16{1500, -32000}
	arrayUint16 := [2]uint16{0, 5000}
	arrayFloat32 := [2]float32{2.5, 1000.50}
	arrayTexto := [2]string{"", "Olá"}
	arrayEstrutura := [2]Pessoa{
		{"Pedro", 10, "Pedrinho"},
		{"Tiago", 5, "Titi"},
	}

	sliceVazioBool := []bool{}
	sliceVazioInt16 := []int16{}
	sliceVazioUint16 := []uint16{}
	sliceVazioFloat32 := []float32{}
	sliceVazioTexto := []string{}
	sliceVazioEstrutura := []Pessoa{}

	sliceBool := []bool{true, false}
	sliceInt16 := []int16{1500, -32000}
	sliceUint16 := []uint16{0, 5000}
	sliceFloat32 := []float32{2.5, 1000.50}
	sliceTexto := []string{"", "Olá"}
	sliceEstrutura := []Pessoa{
		{"Pedro", 10, "Pedrinho"},
		{"Tiago", 5, "Titi"},
	}

	fmt.Println("#### USO DE MÉMORIA (byte = bits / 8) ####")
	imprimir("bool1", unsafe.Sizeof(bool1))
	imprimir("numInt8", unsafe.Sizeof(numInt8))
	imprimir("numInt16", unsafe.Sizeof(numInt16))
	imprimir("numInt32", unsafe.Sizeof(numInt32))
	imprimir("numInt64", unsafe.Sizeof(numInt64))
	imprimir("numUint8", unsafe.Sizeof(numUint8))
	imprimir("numUint16", unsafe.Sizeof(numUint16))
	imprimir("numUint32", unsafe.Sizeof(numUint32))
	imprimir("numUint64", unsafe.Sizeof(numUint64))
	imprimir("realFloat32", unsafe.Sizeof(realFloat32))
	imprimir("texto", unsafe.Sizeof(texto))
	imprimir("estrutura", unsafe.Sizeof(estrutura1))

	fmt.Println("")
	imprimir("arrayBool", unsafe.Sizeof(arrayBool))
	imprimir("arrayInt16", unsafe.Sizeof(arrayInt16))
	imprimir("arrayUint16", unsafe.Sizeof(arrayUint16))
	imprimir("arrayFloat32", unsafe.Sizeof(arrayFloat32))
	imprimir("arrayTexto", unsafe.Sizeof(arrayTexto))
	imprimir("arrayEstrutura", unsafe.Sizeof(arrayEstrutura))
	imprimir("arrayEstrutura0", unsafe.Sizeof(arrayEstrutura[0]))
	imprimir("arrayEstrutura0Nome", unsafe.Sizeof(arrayEstrutura[0].Nome))
	imprimir("arrayEstrutura0Nick", unsafe.Sizeof(arrayEstrutura[0].Nick))
	imprimir("arrayEstrutura0Idade", unsafe.Sizeof(arrayEstrutura[0].Idade))

	fmt.Println("")
	imprimir("sliceVazioBool", unsafe.Sizeof(sliceVazioBool))
	imprimir("sliceVazioInt16", unsafe.Sizeof(sliceVazioInt16))
	imprimir("sliceVazioUint16", unsafe.Sizeof(sliceVazioUint16))
	imprimir("sliceVazioFloat32", unsafe.Sizeof(sliceVazioFloat32))
	imprimir("sliceVazioTexto", unsafe.Sizeof(sliceVazioTexto))
	imprimir("sliceVazioEstrutura", unsafe.Sizeof(sliceVazioEstrutura))

	fmt.Println("")
	imprimir("sliceBool", unsafe.Sizeof(sliceBool))
	imprimir("sliceInt16", unsafe.Sizeof(sliceInt16))
	imprimir("sliceUint16", unsafe.Sizeof(sliceUint16))
	imprimir("sliceFloat32", unsafe.Sizeof(sliceFloat32))
	imprimir("sliceTexto", unsafe.Sizeof(sliceTexto))
	imprimir("sliceEstrutura", unsafe.Sizeof(sliceEstrutura))
	imprimir("sliceEstrutura0", unsafe.Sizeof(sliceEstrutura[0]))
	imprimir("sliceEstrutura0Nome", unsafe.Sizeof(sliceEstrutura[0].Nome))
	imprimir("sliceEstrutura0Nick", unsafe.Sizeof(sliceEstrutura[0].Nick))
	imprimir("sliceEstrutura0Idade", unsafe.Sizeof(sliceEstrutura[0].Idade))
}

func imprimir(nomeVariavel string, totalBytes uintptr) {
	var print string = "byte"
	if totalBytes > 1 {
		print += "s"
	}
	fmt.Printf("%20s\t %02d %s\n", nomeVariavel, totalBytes, print)
}

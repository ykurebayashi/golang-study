package main

import "fmt"

func main() {
	// Declaração de variável
	/* var nomeDaVariável tipagem = valor */
	var nameOne string = "Marcos 1"
	var nameTwo string = "Marcos 2"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	// Reatribuir valor de mesmo tipo a uma variável já declarada
	/* nomeDaVariável = novoValor */
	nameOne = "Marcos novo"
	nameThree = "Marcos 3"
	fmt.Println(nameOne, nameTwo, nameThree)

	// Shorthand para declarar variável dentro de função
	/* := semelhante ao */
	/* A tipagem pode ser inferida quando não é declarada */
	nameFour := "Vasco"
	fmt.Println((nameFour))

	// Outros tipos de variávels
	/* ints */
	var ageOne int = 20
	ageTwo := 40
	fmt.Println(ageOne, ageTwo)

	/* floats */
	var scoreOne float32 = 22.5
	var scoreTwo float32 = 898.2
	fmt.Println(scoreOne, scoreTwo)

	/* bits e memória */
	var numberOne int8 = 25   // int8, de acordo com a doc, é qualquer valor entre -128 e 127
	var numberTwo int16 = 325 // int16, de acordo com a doc, é qualquer valor entre -32768 e 32767
	fmt.Println(numberOne, numberTwo)
}

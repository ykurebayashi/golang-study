package main

import "fmt"

func main() {
	//Print
	fmt.Print("Teste")
	fmt.Print("Teste \n")
	fmt.Print("Nova linha")

	//Print com nova linha
	fmt.Println("Teste")
	fmt.Print("Nova linha \n")

	//Print com variável
	age := 22
	name := "Marcos"
	fmt.Println("Teste de variável, idade:", age, "anos de idade.", "Meu nome é:", name)

	//PrintF (formatted string) %_ = format specifier
	fmt.Printf("Teste de variável, idade: %v anos de idade. Meu nome é: %v \n", age, name)
	fmt.Printf("Teste de variável, idade: %q anos de idade. Meu nome é: %q \n", age, name)
	fmt.Printf("Age é do tipo %T \n", age)
	fmt.Printf("Quantidade de dinheiro: %f \n", 22.4)
	fmt.Printf("Quantidade de dinheiro: %0.1f \n", 22.4)

	//SprintF (save formatted string)
	result := fmt.Sprintf("Teste de variável, idade: %v anos de idade. Meu nome é: %v \n", age, name)
	fmt.Print("Resultado do save formatted string:", result)
}

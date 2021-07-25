// Exercício #01
// Faça um programa em que 3 variáveis recebem valores diferentes e informa qual a variável com maior valor.

package main

import (
	"fmt"
)

func main() {

	Amora := 300
	Morango := 200
	Banana := 100

	if Banana < Amora {
		fmt.Printf("Possuímos mais amoras em estoque e sua quantidade é de %d unidades\n\n", Amora)
	} else if Banana > Amora && Banana > Morango {
		fmt.Printf("Possuímos mais bananas em estoque e sua quantidade é de %d unidades\n\n", Banana)
	} else {
		fmt.Printf("Possuímos mais morangos em estoque e sua quantidade é de %d unidades\n\n", Morango)
	}

	// Exercício #02
	// Faça um programa que testa se um número passado em uma variável é 0, múltiplo de 2, múltiplo de 3 ou nenhuma das opções.

	numero := 25

	if numero%2 == 0 {
		fmt.Printf("O valor é múltiplo de 2")

	} else if numero%3 == 0 {
		fmt.Printf("O valor é múltiplo de 3")

	} else if numero == 0 {
		fmt.Printf("O valor é 0")

	} else {
		fmt.Printf("O valor não é 0 e nem multiplo de 2 ou 3")
	}

}

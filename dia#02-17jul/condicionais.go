// Exemplos de Estudo

// If

package main

import (
	"fmt"
)

func main() {

	idade := 20

	if idade >= 18 {
		fmt.Println("Maior de idade")
	}

	// If else

	if idade < 18 {
		fmt.Println("Menor de idade")
	} else {
		fmt.Println("Maior de idade")
	}

	// If, else if, else

	numero1 := 20
	numero2 := 20

	if numero1 < numero2 {
		fmt.Println("O número 1 é menor que o número 2")
	} else if numero1 > numero2 {
		fmt.Println("O número 1 é maior que o número 2")
	} else {
		fmt.Println("O número 1 e o número 2 são iguais")
	}

	// If com atribuição

	if idade := 20; idade < 18 {
		fmt.Println("Menor de idade")
	} else {
		fmt.Println("Maior de idade")
	}

	//Idiomático em Go
	//Evitar o else e retornar o mais cedo possível

	idade1 := 10
	if idade1 < 18 {
		fmt.Println("Menor que 18 anos, não pode prosseguir")
		return
	}
	fmt.Println("Maior de idade")

}

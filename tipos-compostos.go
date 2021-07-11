// Arrays
// # Exercício 04 - Aula

package main

import (
	"fmt"
)

func main() {

	numeros := [6]string{"zero", "um", "dois", "tres", "quatro", "cinco"}
	fmt.Printf("o tipo é: %T\n", numeros)
	fmt.Println(numeros[:])

	// Slice
	// # Exercício 05 - Aula
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{7, 8, 9, 10, 11, 12}
	slice3 := append(slice1, slice2...)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	// # Exercício 06 - Aula
	signos := make([]string, 12)

	signos[0] = "Áries"
	signos[1] = "Touro"
	signos[2] = "Gêmeos"
	signos[3] = "Câncer"
	signos[4] = "Leão"
	signos[5] = "Virgem"
	signos[6] = "Libra"
	signos[7] = "Escorpião"
	signos[8] = "Sagitário"
	signos[9] = "Capricórnio"
	signos[10] = "Aquário"
	signos[11] = "Peixes"

	fmt.Println(signos)
	fmt.Println(signos[1:4])

	// Maps
	// # Exercício 07 - Aula
	ano := make(map[int]string)
	ano[1] = "Janeiro"
	ano[2] = "Fevereiro"
	ano[3] = "Março"
	ano[4] = "Abril"
	ano[5] = "Maio"
	ano[6] = "Junho"
	ano[7] = "Julho"
	ano[8] = "Agosto"
	ano[9] = "Setembro"
	ano[10] = "Outubro"
	ano[11] = "Novembro"
	ano[12] = "Dezembro"

	fmt.Printf("Mês 1 é %s e mês 12 é %s.\n", ano[1], ano[12])
	fmt.Println(len(ano))
}

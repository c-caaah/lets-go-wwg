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

	// Exercício #01
	// Declare uma variável e atribua a ela o ano de nascimento de uma pessoa.
	// Declare uma variável e atribua a ela o ano atual.
	// Escreva um programa que verifique se no ano atual essa pessoa já está apta a votar e que printe na tela uma mensagem informando caso positivo e caso negativo.

	AnoNasc := 1991
	AnoAtual := 2021
	idadeVoto := AnoAtual - AnoNasc

	if idadeVoto >= 16 {
		fmt.Printf("O usuário possui %d anos e está apto para votar\n", idadeVoto)
		return
	}
	fmt.Printf("O usuário possui %d anos e não está apto para votar\n", idadeVoto)

	// Exercício #02
	// Declara uma variável;
	// Atribui o valor de um número a ela;
	// Informa se o número é positivo ou negativo.

	NumAleatorio := -35

	if NumAleatorio > 0 {
		fmt.Println("o número delarado é positivo")
	} else if NumAleatorio < 0 {
		fmt.Println("o número declarado é negativo")
	} else {
		fmt.Println("o número declarado é 0")
	}

	// Exercício #03
	// Faça um programa que, dada a idade de uma pessoa, verifique se ela é menor de idade, maior de idade ou idosa utilizando a instrução if.
	// Leve em consideração os seguintes valores:
	// Faixa: Menor de idade, Maior de idade, Idoso
	// Intervalos: abaixo de 18, entre 18 e 60, acima de 60

	idade := 30

	if idadeUsuario >= 60 {
		fmt.Printf("O usuário possui %d anos e é idoso", idadeUsuario)
	} else if idadeUsuario >= 18 && idadeUsuario < 60 {
		fmt.Printf("O usuário possui %d anos e é maior de idade", idadeUsuario)
	} else if idadeUsuario < 0 {
		fmt.Println("Idade inválida")
	} else {
		fmt.Printf("O usuário possui %d anos e é menor de idade", idadeUsuario)
	}

	// Exercício #04
	// Faça um programa a partir das mesmas orientações do exercício acima, mas utilizando switch em vez de if.

	// <em breve>

	// Exercício #05
	// Declare uma variável chamada hora e atribuir um valor int a ela.
	// Usando switch, descreva cases e printe na tela se a hora corresponde ao período da manhã, tarde, noite ou madrugada.

	// <em breve>
}

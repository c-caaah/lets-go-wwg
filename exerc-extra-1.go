// A aplicação abaixo não compila
// Descubra por que não compila
// Erros de compilação nos ajudam a compreender o que precisamos consertar em nosso código. O que o erro ./prog.go:9:14: constant 150 overflows int8 nos indica?
// - o valor é maior que o valor máximo do int
// Conserte o erro de compilação e faça a quantidade de quilômetros ser imprimida na tela
// - precisa estar entre -128 ~ 127

package main

import (
	"fmt"
)

func main() {
	var quilometros int8
	quilometros = 126

	fmt.Println(quilometros)
}

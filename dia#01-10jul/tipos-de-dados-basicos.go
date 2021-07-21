// Ponto Flutuante
// # Exercício 01 - Aula
package main

import (
	"fmt"
)

func main() {
	var pesoDoTomate, quantidadeDeGarrafasDeAzeite, unidadesDeSabonete float64
	pesoDoTomate = 2.600
	quantidadeDeGarrafasDeAzeite = 2
	unidadesDeSabonete = 7

	var precoDoTomate, precoDoAzeite, precoDoSabonete float64
	precoDoTomate = 10.3
	precoDoAzeite = 19
	precoDoSabonete = 5.80

	valorDaCompra := (pesoDoTomate * precoDoTomate) + (quantidadeDeGarrafasDeAzeite * precoDoAzeite) + (unidadesDeSabonete * precoDoSabonete)
	fmt.Printf(">> o valor da compra deu: %v \nSó uma garrafa de azeite já custou %v", valorDaCompra, precoDoAzeite)

	//String
	// # Exercício 02 - Aula
	nome := "Cáh"
	cor := "Roxa"

	fmt.Printf("\n\n>> Olá, meu nome é %s e eu gosto da cor %s.", nome, cor)

	// Booleano
	// # Exercício 03 - Aula
	a := 15 >= 15         // true
	b := 100 < 1000       // true
	c := 10 == 18         // true
	d := 5 != 5           // false
	e := 89 > 0 && 50 > 0 // true

	fmt.Printf("\n\no tipo de a é %T e seu valor é %v\n", a, a)
	fmt.Printf("o tipo de b é %T e seu valor é %v\n", b, b)
	fmt.Printf("o tipo de c é %T e seu valor é %v\n", c, c)
	fmt.Printf("o tipo de d é %T e seu valor é %v\n", d, d)
	fmt.Printf("o tipo de e é %T e seu valor é %v\n", e, e)

}

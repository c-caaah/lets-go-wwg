// Exercício Extra 01 - Arrays
// Existem dois times de futebol, o time amarelo e o time vermelho. O time amarelo tem 5 jogadores (Fernando, João, Lúcia, Mariana e Ana) e o time vermelho tem 4 jogadores (Helena, Jonas, José e Juliana).
// Crie um array de string para cada time e nomeie com o nome do time.
// Printe na tela os nomes dos jogadores de cada time

package main

import (
	"fmt"
)

func main() {

	TimeAmarelo := [5]string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}
	TimeVermelho := [4]string{"Helena", "Jonas", "José", "Juliana"}

	fmt.Printf("O time Amarelo é composto por %s\n", TimeAmarelo)
	fmt.Printf("O time Vermelho é composto por %s\n\n", TimeVermelho)

	// Exercício 02 - Slice	usando append()
	// Considerando os times do item anterior, crie uma slice para representar cada um.
	// Adicione o jogador Luis Inácio no time vermelho usando a função append()
	// Printe na tela os nomes dos jogadores do time vermelho

	Amarelo := []string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}
	Vermelho := []string{"Helena", "Jonas", "José", "Juliana"}

	Vermelho = append(Vermelho, "Luis Inácio")

	fmt.Printf("O time amarelo é composto por %s\n", Amarelo)
	fmt.Printf("O time vermelho é composto por %s\n\n", Vermelho)

	// Exercício 03 - Maps
	// Escreva um programa que lista o nome dos países e quantas vezes eles aparecem no nosso map.
	// Passos:
	// Criar um mapa com chave do tipo string e valor do tipo string (map[string]string) onde a chave seja o nome da cidade e o valor o nome do país.
	// Percorrer o mapa do item 1 acumulando em outro mapa a frequência de aparições do país. Esse mapa segundo mapa terá tipo map[string]int, sendo a chave o nome do país e o valor a quantidade de menções.
	// Printe na tela os valores do último mapa.

	Paises := map[string]string{
		"Munique":           "Alemanha",
		"Frankfurt am main": "Alemanha",
		"Galway":            "Irlanda",
		"Glasgow":           "Escócia",
		"Liverpool":         "Inglaterra",
		"Curitiba":          "Brasil",
		"Waterford":         "Irlanda",
	}

	QtdePaises := make(map[string]int)
	for _, pais := range Paises {
		QtdePaises[pais]++
	}

	fmt.Printf("Contagem de citação dos Países:\n%v\n", QtdePaises)

}

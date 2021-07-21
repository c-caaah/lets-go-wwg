// Continuação do dia 10/07
// Structs

// Exemplo 1
package main

import (
	"fmt"
)

func main() {

	type Pessoa struct {
		Nome  string
		Idade int
		Cor   string
	}
	pessoa := Pessoa{
		Nome:  "Marta",
		Idade: 20,
	}
	fmt.Printf("%+v\n", pessoa)

	// Acessando os elementos de pessoas
	fmt.Printf("%s tem %d anos\n", pessoa.Nome, pessoa.Idade)

	// Exemplo 2

	type Aula struct {
		Descricao string
		Docente   Pessoa
		Discente  []Pessoa
	}

	eliana := Pessoa{
		Nome:  "Eliana",
		Idade: 47,
	}

	joao := Pessoa{
		Nome: "João", 
		Idade: 19
	}

	zeca := Pessoa{
		Nome: "Zeca", 
		Idade: 34
	}

	aline := Pessoa{
		Nome: "Aline", 
		Idade: 42
	}

	aula := Aula{
		Descricao: "Aula de Go",
		Docente:   eliana,
		Discente: []Pessoa{
			joao,
			zeca,
			aline,
		},
	}

	fmt.Printf("%+v", aula)

	// Exercício 1
	// Crie um tipo Pessoa que tenha os atributos nome, idade e cor preferida.
	// Crie três variáveis do tipo pessoa.
	// Printe o nome de todas as três, bem como frases informando sua idade e cores preferidas.

	dionisio := Pessoa{
		Nome: "Dionísio", 
		Idade: 25, 
		Cor: "Amarelo"
	}

	gertrudes := Pessoa{
		Nome: "Gertrudes", 
		Idade: 33, 
		Cor: "Preto"
	}

	abighail := Pessoa{
		Nome: "Abighail", 
		Idade: 55, 
		Cor: "Azul"
	}

	fmt.Printf("%s tem %d anos e sua cor preferida é %s\n", dionisio.Nome, dionisio.Idade, dionisio.Cor)
	fmt.Printf("%s tem %d anos e sua cor preferida é %s\n", gertrudes.Nome, gertrudes.Idade, gertrudes.Cor)
	fmt.Printf("%s tem %d anos e sua cor preferida é %s", abighail.Nome, abighail.Idade, abighail.Cor)

}

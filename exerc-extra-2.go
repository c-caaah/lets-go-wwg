// Dado o ano em que a pessoa nasceu, calcule quantos anos ela tem no ano atual (para fins de arredondamento, considere que ela já fez aniversário no ano atual).
// Como podemos pegar a informação do ano diretamente do console?
// Dicas:
// O método time.Now() retorna a data atual
// Para ler do console, você pode utilizar o método fmt.Scan()

package main

import (
	"fmt"
	"time"
)

func main() {

	AnoAtual := time.Now().Year()
	var AnoNasc int
	fmt.Println("Ano de Nascimento: ")
	fmt.Scanf("%d", &AnoNasc)
	idade := AnoAtual - AnoNasc
	fmt.Printf("Você tem %d anos.", idade)

}

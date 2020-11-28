package main

import "fmt"

type moedas struct {
	moeda string
	id    int
}

// Maps em Go, são parecidos como Maps em Python, recebem uma key e values, em pares
func main() {
	// Sintaxe de criação de um map, passamo o tipo da key, dentro dos colchetes []
	// e o tipo do value logo em seguida
	estadosCapitais := make(map[string]string)
	fmt.Println(estadosCapitais)

	// Para adicionar valores, podemos passar diretamente, a key e o value
	estadosCapitais["RS"] = "Porto Alegre"
	estadosCapitais["SP"] = "São Paulo"
	estadosCapitais["SC"] = "Florianopolis"
	fmt.Println("Estados e capitais:", estadosCapitais)

	// Podemos também definir os dados no momento da sua criação
	produtos := map[string]int{
		"Sabão em pó": 10,
		"Água":        2,
		"Maçã":        1,
	}

	fmt.Println("Map de produtos: ", produtos)
	// As keys do map podem ser qualquer tipo de dado comparável
	// (bool, inteiro, float, complex, string)
	// As keys podem ser até strutcs, definidas pelo user.
	// Maps iniciam com valor nil, caso não seja passado nenhum value ou key
	// Podemos pegar o valor passando uma key, tanto em uma variável com o mesmo tipo da key, como
	// o tipo da key direto
	nomeProduto := "Água"
	produto := produtos[nomeProduto]
	fmt.Println("Produto", nomeProduto, "custa", produto, "bonoros")

	// Podemos verificar se uma key existe
	value, ok := produtos["Nutellinha"]
	// Com a sintaxe acima, o value recebe o valor se a key exister,
	// e o ok recebe true se a key existir
	if ok == true {
		fmt.Println("Produto custa", value)
	}
	fmt.Println("Produto não existe", ok)

	// Podemos iterar pelos valores usando o for range, onde teremos o valor da key e o value
	for key, value := range produtos {
		fmt.Println("Produto -", key, "- valor -", value)
	}

	// Para remover itens do map, podemos usar a função delete.
	// Nela, passamos o map e a key que queremos remover
	delete(produtos, "Maçã")
	fmt.Println("após remoção", produtos)

	// Podemos usar o len para ver o tamanho do map
	fmt.Println("Tamanho do map é", len(produtos))

	// Maps também podem conter dados do tipo struct
	s1 := moedas{
		moeda: "R$",
		id:    1,
	}

	s2 := moedas{
		moeda: "$",
		id:    2,
	}

	infoMoedas := map[string]moedas{
		"Brasi": s1,
		"EUA":   s2,
	}

	for pais, info := range infoMoedas {
		fmt.Printf("Pais: %s, Cifrão: %s, ID: %d\n", pais, info.moeda, info.id)
	}

	// Quando pegamos um valor com a key, ele funciona da mesma forma que um slice,
	// qualquer alteração feita no dado, irá refletir no map original.

	// Maps não podem ser comparados entre eles, mesmo que identicos,
	// eles só podem ser comparados com nil

}

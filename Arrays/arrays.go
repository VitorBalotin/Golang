package main

import "fmt"

// Um Array é uma coleção de elementos que são do mesmo tipo.
// Não é possivel misturar tipos de dados em Go.
func main() {
	// Definição de array [n]T, com var, onde [n] é o tamanho e T é o tipo
	var a [2]string

	// outra forma de definir array, com :=
	numeros := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numeros)

	// Passando parametros para o array
	a[0] = "Hello"
	a[1] = "World"

	// range usando index e value
	// podemos também usar o _ (blank identifier, identificador vazio) no lugar
	// do index se quisermos ignorar ele
	for _, value := range a {
		fmt.Print(value, " ")
	}

	fmt.Printf("\n")

	// *******************************************************************************

	// Apos ser definido dessa forma, onde o compilador diz o seu tamanho,
	// não é possivel inserir mais valores
	b := [...]int{1, 2, 3, 4}
	fmt.Println("Array definido ", b)

	// O tamanho do array é visto com a função len(array)
	// *******************************************************************************
	fmt.Println("Tamanho do array", len(b))

	fmt.Println("Array definido ", b)

	fmt.Println("Print de cada numero passando por um for")
	for i := range numeros {
		// Printf, funciona parecido como o do C, onde %s é string e %d é para numeros,
		// que irão substituir a frase
		fmt.Printf("O numero atual é %d.\n", numeros[i])
	}

	// Existem arrays multidimensionais, parecidos com matrizes
	c := [3][2]string{
		{"1,1", "1,2"},
		{"2,1", "2,2"},
		{"3,1", "3,2"}, // compilador chora se não tiver essa virgula ali
	}

	fmt.Println(c)
}

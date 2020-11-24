package main

import "fmt"

func main() {
	// Definição de array, com var
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println("Print da primeira posição do array.")
	fmt.Println(a[0])
	fmt.Println("Print da segunda posição do array.")
	fmt.Println(a[1])

	fmt.Println("Print do array completo passando cada parametro.")
	fmt.Println(a[0], a[1])

	fmt.Println("Print do array completo passando passando apenas o array.")
	fmt.Println(a)

	fmt.Println("Print de cada item passando por um for")
	for i := range a {
		fmt.Print(a[i], " ")
	}

	fmt.Printf("\n")

	// outra forma de definir array, com :=
	numeros := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numeros)

	fmt.Println("Print de cada numero passando por um for")
	for i := range numeros {
		// Printf, funciona parecido como o do C, onde %s é string e %d é para numeros,
		// que irão substituir a frase
		fmt.Printf("O numero atual é %d.\n", numeros[i])
	}


	// Slices tem tamanho dinamico e oferecem uma view flexivel dos elementos de um array
	// passa-se o [indice inicial : indice final] onde : é o separador.
	var slice []int = numeros[1:3]

	fmt.Println("Slice", slice)
}

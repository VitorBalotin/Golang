package main

import "fmt"

func diminuiUm(numbers []int) {
	for i := range numbers {
		numbers[i]--
	}
}

func paises() []string {
	paises := []string{"Alemanha", "Brasil", "Canadá", "Japão"}
	slicePaises := paises[:len(paises)-2]
	paisesCopy := make([]string, len(slicePaises))
	copy(paisesCopy, slicePaises)
	return paisesCopy
}

// Slices são referencias para um array que já existe,
// funciona como um wrapper, por cima do array já definido.
// Não possui nenhum dos dados.
func main() {
	// Definindo um slice
	// array que vamos referenciar
	a := [5]int{20, 30, 40, 50, 60}
	// a sintaxe é array[inicio : fim - 1], ou pode ser passado sem os parametros,
	// para referenciar o array todo
	var b []int = a[2:5]
	fmt.Println(b)

	// Aqui estamos definindo um slice c que vai referenciar um array com 3 inteiros
	c := []int{2, 3, 4}
	fmt.Println(c)

	// Aqui vamos ver como as alterações refletem do slice para o array,
	// como o slice e o array são basicamente os mesmos dados, qualquer alteração no slice
	// reflete diretamente no array.
	d := [...]int{23, 39, 15, 25, 75, 90, 128}
	dSlice := d[3:7]
	fmt.Println("array antes da operação", d)
	for i := range dSlice {
		dSlice[i]++
	}
	fmt.Println("array após da operação", d)

	// Aqui vamos ver bem como os dados são compartilhados
	// Vemos também uma nova forma de definir
	arrayNumbers := [3]int{90, 91, 92}
	sliceNumbers1 := arrayNumbers[:]
	sliceNumbers2 := arrayNumbers[:]
	fmt.Println("array original", arrayNumbers)
	// Passamos um novo valor na posição [0] no slice1
	sliceNumbers1[0] = 10
	fmt.Println("array ao modificar o item 0 no sliceNumbers1", arrayNumbers)
	// Agora alteraremos a posição [0] no slice2
	sliceNumbers2[0] = 20
	fmt.Println("array ao modificar o item 0 no sliceNumbers2", arrayNumbers)

	// Aqui veremos o tamanho que o array e o slice terão e capacidade do slice
	arrayFruit := [...]string{"maçã", "pera", "tomate", "laranja"}
	sliceFruit := arrayFruit[1:3]
	fmt.Println("Tamanho do slice", len(arrayFruit), ", capacidade do slice", cap(sliceFruit))

	// O slice pode ser re sliced, mas somente até sua capacidade
	sliceFruit = sliceFruit[:cap(sliceFruit)]
	fmt.Println("Após ser feito o slice novamente, o tamanho é", len(sliceFruit), "e a capacidade é ", cap(sliceFruit))

	// Slices podem ser criados com a função make([]T, tamanho, capacidade)
	sliceMake := make([]int, 5, 5)
	// o slice é criado zerado
	fmt.Println(sliceMake)

	// Como slices são dinamicos, podemos inserir mais que a capacidade usando a função append(),
	// que é uma função variadic, sua definição é append(s []T, x ...T) []T.
	// O que o append faz, é criar um novo array por baixo dos panos, duplicando a capacidade e
	// os elementos do array original é copiado para o novo

	cars := []string{"Honda", "Ford", "Chevrolet"}
	fmt.Println("Carros: ", cars, "tem tamanho", len(cars), "e capacidade", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("Carros:", cars, "novo tamanho", len(cars), "e nova capacidade", cap(cars))

	// iniciamos um slice vazio, que é nill
	var nomes []string
	fmt.Println(nomes)
	nomes = append(nomes, "Thai", "Vito", "Léo", "Welito")
	fmt.Println("Nomes:", nomes)

	// Podemos usar o operador ... para criar um novo slice com o append
	vegetais := []string{"batata", "berinjela"}
	frutas := []string{"maçã", "pera", "tomate", "laranja"}
	comida := append(vegetais, frutas...)
	fmt.Println("comida:", comida)

	// Se passarmos o slice para uma função, por ele conter suas informações de capacidade,
	// tamanho e ponteiro para o elemento original do array, quando passado por parametro,
	// para uma função, criar função acima.
	numeros := []int{2, 3, 4}
	fmt.Println("Slice antes da chamada da função", numeros)
	diminuiUm(numeros)
	fmt.Println("Slice após a chamada da função", numeros)

	// Slices também podem ser multidimensionais
	lps := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Swift"}, // compilador chora sem a virgula =(
	}
	fmt.Println(lps)

	// Bom, como slices seguram uma referencia para um array, enquanto o slice estiver em memoria
	// ele não pode ser removido pela garbage collection, isso pode ser um problema, já que nem sempre
	// vamos utilizar o array todo, o que pode ser feito é utilizar a função copy

	slicePaises := paises()
	fmt.Println(slicePaises)
}

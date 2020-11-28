package main

import "fmt"

func altera(val *int) {
	*val = 12
}

func returnPointer() *int {
	i := 20
	return &i
}

func arrayMod(arr *[3]int) {
	(*arr)[0] = 90
}

// Ponteiros em Go se difere de linguagens como C e C++
// Ponteiros gravam o endereço de memória de uma outra variável
func main() {
	// Aqui vamos apontar e mostrar o endereço de memória
	b := "apontado"
	// Para criar o ponteiro a temos que passar o tipo da variável que vamos apontar usando o *
	// no tipo do dado e o & para pegar o endereço de memória.
	// O * serve também para pegar o valor de dentro do ponteiro. Chamado de Dereferencing
	var a *string = &b

	fmt.Printf("O tipo de a é %T\n", a)
	fmt.Println("O endereço de b é", a)

	// O valor do ponteiro é NIL, quando iniciado
	c := 10
	var d *int
	fmt.Println("Valor inicial do ponteiro", d)

	if d == nil {
		// Aqui, apontamos para c usando o operador &
		d = &c
		fmt.Println("b após ser iniciado", d)
	}

	// O Go também tem a função new, para criação de ponteiros
	// Essa função recebe o tipo do dado, como argumento e retorna um endereço que ela alocará
	// na memória

	ponteiro := new(int)

	fmt.Printf("O valor do ponteiro é %d, o tipo é %T, e o endereço é %v\n",
		*ponteiro, ponteiro, ponteiro)
	// Vamos alocar um valor
	*ponteiro = 20
	fmt.Println("Novo valor é", *ponteiro)

	// Vamos fazer uma operação/alteração em cima de um ponteiro
	e := 255
	f := &e
	fmt.Println("O endereço de f é", f)
	fmt.Println("Valor de f é", *f)
	*f++
	fmt.Println("O novo valor de e é", e)

	// Aqui nós passamos o ponteiro para um função que espera receber um ponteiro
	// e vai alterar o valor original
	valueFunc := 90
	fmt.Println("Valor antes da chamada da função", valueFunc)
	pointerFunc := &valueFunc
	altera(pointerFunc)
	fmt.Println("Valor depois da chamada da função", valueFunc)

	// Aqui a função vai retornar um ponteiro para algo definido acima
	pointerReturn := returnPointer()
	fmt.Println("Valor do retorno da função", *pointerReturn)

	// Aqui nós vamos criar um ponteiro para um array, isso é totalmente desnecessário,
	// pois temos slices, mas, é possivel criar também
	arrayzao := [3]int{10, 20, 30}
	arrayMod(&arrayzao)
	fmt.Println(arrayzao)

	// Go não tem suporte para aritmética de ponteiro, como C e C++
	arrArith := [...]int{109, 110, 111}
	pointerArith := &arrArith
	// Comentado porque ele vai jogar um erro, operação invalida
	// pointerArith++
	fmt.Println(pointerArith)
}

package main

import "fmt"

func main() {
	sum := 0

	//	---------------------FOR BÁSICO-------------------------

	// Go só tem um tipo de loop, que é o For, mas ele pode ser usado de várias maneiras
	// Num for básico em golang temos três components separados por ponto e virgula:

	//   	O argumento inicial que é executado antes da primeira iteração
	//      O argumento inicial geralmente vai ser uma declaração de variavel que vai ficar somente dentro do escopo do if, mas ele nao é um argumento obrigaótior

	//		A expressão de condição que vai ser avaliada antes de cada iteração, se a condição deixar de ser verdadeira o loop vai parar
	//      Ao contrário de outras linguagens como C, Java e JS em GO não existem parenteses ao redor da condição do for mas as chaves { } são obrigatórias

	//		E o argumento final que vai ser executado no fim de cada iteração
	//      O argumento final também não é obrigatório

	fmt.Println("\n----FOR BÁSICO----")
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println("A soma no turno", i, "é:", sum)
	}

	//	--------- FOR BÁSICO USANDO BREAK----------------------------

	//Num for também podemos usar um break pra parar a iteração antes de cumprir a condição expressada

	fmt.Println("\n---FOR BÁSICO COM BREAK---")
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d", i)
	}

	// ------------FOR BÁSICO COM CONTINUE---------------------------

	// for básico usando o argumento continue
	// ele é usado pra pular a iteração atual do loop
	// todo o código presente dentro do for no loop nao vai ser executado durante a iteração atual
	// e o loop vai prosseguir pra proxima iteração

	//Checa o resto de divisão por dois, se for zero então o número é par e o argumento continue vai ser executado
	//forçando o for a ir pra próxima iteração
	//então o output vai ser apenas dos números impares do for

	fmt.Println("\n\n-----FOR BÁSICO COM CONTINUE-----")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}

	//-----------------------FOR COMO WHILE-------------------------

	//No caso de um WHILE em GO ele é declarado como um FOR, não vamos precisar utilizar ponto e virgula para separar argumentos

	fmt.Println("\n\n------WHILE COMO FOR------")
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("Total da soma:", sum)

	//---------------------FOR COMO LOOP INFINITO------------------------

	// Um For em GO pode ser um loop infinito, só é necessário declarar ele sem condição e sem parenteses

	fmt.Println("\n----FOR INFINITO----")
	fmt.Println("---Esse for seria infinito,mas tem um break dentro---")
	value := 0
	for {
		value = value + 1
		if value > 5 {
			fmt.Println("Chegou no break point, o valor é: ", value)
			break
		}
	}

	//-------------------FOR USANDO RANGE---------------------------------

	// o range faz o FOR iterar elementos dentro de uma estrutura de dados
	// entao criamos um array e usamos esse array pra fazer a iteração

	// For usando range, iterando por um array
	// Pra usar range com arrays ou slices nós precisamos prover o for com o index e o valor pra cada entrada.
	// index é o i e o valor é o numbers[i]

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("\n----FOR usando RANGE pra iterar os  items dentro do array de numeros----")
	for i := range numbers {
		fmt.Println("Posição é:", i, "e o numero nessa posição do array é:", numbers[i])
	}

	//RANGE num map itera usando uma chave ligada a um valor
	fmt.Println("\n----FOR usando RANGE e as chaves do map pra iterar----")
	capitalCountrys := map[string]string{"Alemanha": "Berlin", "Irlanda": "Dublin", "Italia": "Roma"}
	for country := range capitalCountrys {
		fmt.Println("Capital da(e)", country, "é", capitalCountrys[country])
	}

	//Aqui é a mesma lógica, mas passamos country e capital como key-value no inicio do for
	//e ele pega por padrão dentro do map o primeiro como chave e o segundo como valor
	fmt.Println("\n----FOR usando RANGE e o key-value para iterar----")
	for country, capital := range capitalCountrys {
		fmt.Println("Capital da", country, "é", capital)
	}

	//-------------------FOR EM LOOPS ANINHADOS------------------------------

	// Um FOR que tem outro FOR dentro é chamado de LOOP ANINHADO
	// um loop dentro de outro basicamente

	//a variavel n n vai guardar o numero de linhas na sequencia, nesse caso vamos usar 5
	// o loop de fora vai iterar i de 0 a 4
	//e o loop interno vai iterar j de 0 ao valor atual de i
	//o loop interno vai printar um asterisco pra cada iteração e o loop de fora vai printar uma nova linha a cada iteração
	fmt.Println("\n----FOR EM LOOP ANINHADO----")
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

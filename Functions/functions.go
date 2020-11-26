package main

import (
	"fmt"
)

//Uma função é um bloco de código que realiza uma tarefa especifica
//Ela recebe um input, pode fazer algum cálculo com esse input e gerar um output

//A declaração de uma função começa com a keyword func seguida pelo nome que vai ser chamada a função
//Os parametros são especificados dentro de parenteses seguido pelo tipo de retorno esperado da função
//A sintaxe para especificar um parametro é o nome do parametro primeiro e depois o tipo dele
//Podem ser especificados quantos parametros forem necessários (parameter1 type, parameter2 type)
//Em seguida temos o bloco de código dentro de chaves {} que é o corpo da função

//-------------EXEMPLO DE FUNÇÃO SEM ESPECIFICAR O TIPO DE PARAMETRO E RETORNO----------

//Os tipos de parametro e retorno são opcionais numa função.
//Por isso essa função sem parametro e sem retorno é uma sintaxe válida tambem
func functionname() {
}

//-------------EXEMPLO DE FUNÇÃO COM DOIS PARAMETROS DO MESMO TIPO---------------------

// ao invés de declararmos na função (price int, no int) se eles forem do mesmo tipo
// podemos declarar o tipo apenas uma vez no ultimo parametro
//e entao todos os parametros vao ser desse tipo por padrao

//A funçaõ do exemplo abaixo tem dois parametros de input do mesmo tipo
// e tem como retorno precoTotal que é o produto da operação entre price e no
// entao o valor de retorno tambem é um int

func calculaConta(price, no int) int {
	var precoTotal = price * no
	return precoTotal
}

//--------------EXEMPLO DE FUNÇÃO COM MULTIPLOS VALORES DE RETORNO--------------------

//É possivel retornar multiplos valores de uma função
//Se uma função tem mais de um valor pra retornar os tipos tem que estar especificados dentro de parentese

//Essa função tem como input o comprimento e a largura de um retangula
//e retorna a area e o perimeto dele
//A area do retangulo é o produto do comprimeto e largura
//E o perimeto é duas vezes a soma do comprimento e largura
//todas as variaveis sao float

func areaRetangulo(comprimento, largura float64) (float64, float64) {
	var area = comprimento * largura
	var perimetro = (comprimento + largura) * 2
	return area, perimetro
}

//--------------EXEMPLO DE FUNÇÕES COM VALORES DE RETORNO NOMEADOS-------------------

//é possivel retornar um valor nomeado de uma função. Se um valor de retorno é nomeado, ele pode ser considerado como sendo declarado como uma varaivel na primeira
//linha da função, no caso posso reescrever a função de cima sem especificar o retorno de area e perimetro
//porque eles ja estão nomeados direto no retorno

func rectProps(length, width float64) (area2, perimeter float64) { //aqui eles estao nomeados junto com o tipo esperado de retorno
	area2 = length * width
	perimeter = (length + width) * 2
	return //nao especifica o valor de retorno
}

//-------------------MAIN PARA CHAMADA DE FUNC-------------------

func main() {

	// chamamos a função que a gente criou no método main passando os valores 90 e 6 como parametro 1 e parametro 2 em ordem
	// a função vai fazer o cálculo e retornar o resultado, q vai ser guardado na variavel precoTotal
	fmt.Println("\n---CHAMADA DA FUNC CALCULA CONTA---")
	precoTotal := calculaConta(90, 6)
	fmt.Println("Preco total é:", precoTotal)

	//chamada da função area de um retangulo
	fmt.Println("\n---CHAMADA DA FUNC AREA DE UM RETANGULO---")
	area, perimetro := areaRetangulo(10.8, 5.6)
	fmt.Printf("Area: %f | Perimetro: %f", area, perimetro)

	//----------------EXEMPLO DE FUNÇÃO COM IDENTIFICADOR EM BRANCO---------------------

	//chamada da função rectProps com um identificador em branco
	//_ (underline) é chamado de indentificado em branco em GO ou blank identifier
	//ele é usado no lugar de qualquer valor ou qualquer tipo
	//no caso do exemplo do retangulo a gente usa o underline
	//pra usar apenas o retorno do calculo de area e nao o de perimetro
	//dessa forma o segundo parametro é descartado

	fmt.Println("\n\n---CHAMADA DA FUNC COM UM IDENTIFIER EM BRANCO---")
	area2, _ := rectProps(10.8, 5.6) // perimeter is discarded
	fmt.Printf("Area %f ", area2)
	fmt.Println("\n")

}

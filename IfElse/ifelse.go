package main

import (
	"fmt"
)

func main() {

	//new intn gera apenas um numero aleatorio e é sempre o mesmo, por isso usa o rand.new source que gera um
	//numero aleatorio novo a toda execução
	// x1 := rand.NewSource(time.Now().UnixNano())
	// y1 := rand.New(x1)
	// a := y1.Intn(15 - 0)
	// fmt.Println(a)

	a := 5

	//uma das diferenças do if em GO é que não precisamos colocar a expressao de condição dentro de parenteses
	//mas as chaves sao obrigatórias

	if a > 10 {
		fmt.Println("A é maior que 10")
	} else if a > 5 {
		fmt.Println("A é maior que 5 e menor que 10")
	} else { //detalhe do GO é que nao pode quebrar a linha de comando do if else que dá um erro de sintaxe
		fmt.Println("A é menor ou igual a 5")
	}

	//detalhe do go: pode definir a variavel no proprio if
	//isso vai fazer com que a variavel b fique disponivel somente dentro dos blocos, if, else if e else
	if b := 11; b > 10 {
		fmt.Println("B é maior que 10")
	}

}

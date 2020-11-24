// Linguagem Go
// Integrantes -
// Leonardo Hollas  - 165712
// Thainá Caglioni  - 158286
// Vitor Balotin    - 163252
// Weliton da Silva - 166404

// Aqui criamos o nosso módulo
// Vamos precisar rodar mais um comando, pois temos apenas um package no momento, não um módulo.
// O comando que precisaremos usar é o 'go mod init filename.go ', assim teremos o módulo inicializado
// e vamos conseguir utiliza-lo usando o 'package hello'

// Nosso módulo contém uma dependencia quote, o import é feito
// automaticamente, caso a dependencia não seja encontrada no nosso go.mod,
// mas, agora nosso módulos está totalmente dependente dessa dependencia.
// Essa dependencia que usamos pode trazer outras em conjunto, podemos usar o comando 'go list -m all'
// para visualizarmos o restante delas.
// Caso precise atualizar alguma das dependencias, podemos usar o comando go get 'nome_dependencia'.
// Se precisarmos usar uma versão especifica da dependencia, utilizamos o mesmo comando,
// mas colocamos @vx.x.x.

// Package hello retorna duas funções.
package hello

import "rsc.io/quote"

// Hello retorna uma mensagem de Hello world no idioma do seu sistema operacional.
func Hello() string {
	return quote.Hello()
}

// Go retorna uma mensagem padrão
func Go() string {
	return quote.Go()
}

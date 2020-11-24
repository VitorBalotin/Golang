package main

import (
	"encoding/json"
	"net/http"
)

// Pessoa blabla
type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

// Pessoas blabla
var Pessoas []Pessoa

func buscarTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Pessoas)
}

func buscarPeloNome() {
}

func buscarPorID() {
}

func atualizar() {
}

func cadastrar() {
}

func deletar() {
}

func requestHandler() {
	// Cria a instancia do mux
	// router := mux.NewRouter().StrictSlash(true)

	// pode ser usado o http.HandleFunc
	http.HandleFunc("/pessoas", buscarTodos)
	// router.HandleFunc("/outro", func)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	Pessoas = []Pessoa{
		Pessoa{Nome: "Teste", Idade: 20},
		Pessoa{Nome: "Teste 2", Idade: 40},
	}
	requestHandler()
}

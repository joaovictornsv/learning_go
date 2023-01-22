package main

import "fmt"

type Pessoa struct {
	nome  string
	idade uint8
}

type Estudante struct {
	Pessoa
	curso     string
	faculdade string
}

func main() {
	var pessoa = Pessoa{
		nome:  "João Victor",
		idade: 21,
	}

	var estudante = Estudante{
		Pessoa:    pessoa,
		curso:     "Engenharia da Computação",
		faculdade: "IFPB",
	}

	fmt.Println(estudante)
}

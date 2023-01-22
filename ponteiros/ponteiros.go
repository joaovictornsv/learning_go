package main

import "fmt"

func main() {
	var numero int = 10
	var ponteiro_numero *int = &numero

	fmt.Println(numero, *ponteiro_numero, ponteiro_numero)

	numero = 20

	fmt.Println(numero, *ponteiro_numero, ponteiro_numero)
}

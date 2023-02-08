package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	go escreverVariasVezes("Texto da goroutine", channel1) // goroutine

	fmt.Println("Print antes da função escrever começar de fato")

	// FORMA 1 - Loop infinito, que ficará ouvindo os recebimentos do canal
	// Caso o canal esteja aberto sem receber mensagens ocorrerá um erro de deadlock
	for {

		// Aqui o programa fica "suspenso" aguardando o canal receber algo, quando receber ele continua a execução
		message, is_open := <-channel1
		if !is_open { // Verifica se o canal ainda está aberto
			break
		}
		fmt.Println("Mensagem do canal:", message)
	}

	// FORMA 2 - Usando range, para percorrer cada mensagem do canal, enquanto esse está aberto
	channel2 := make(chan string)
	go escreverVariasVezes("Texto da goroutine", channel2) // goroutine

	for message := range channel2 {
		fmt.Println("Mensagem do canal:", message)
	}

	fmt.Println("Fim do programa")
}

func escreverVariasVezes(texto string, ch chan string) {
	time.Sleep(time.Second * 2) // Delay antes de iniciar a função

	for i := 0; i < 5; i++ {
		ch <- texto // Enviando messagem para o canal
		time.Sleep(time.Second)
	}

	close(ch) // Fechando o canal
}

package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := escrever("Canal 1", 2000)
	channel2 := escrever("Canal 2", 1000)

	exitChannel := multiplex(channel1, channel2)

	for i := 0; i < 10; i++ {
		message := <-exitChannel
		fmt.Println("Mensagem recebida:", message)
	}

}

// Juntando dois canais em um só
// Ou seja, irei ouvir mensagens de dois canais através de um
func multiplex(channel1, channel2 <-chan string) <-chan string {
	exitChannel := make(chan string)

	// Repassando as mensagens recebidas para o canal de saída
	go func() {
		for {
			select {
			case message := <-channel1:
				exitChannel <- message
			case message := <-channel2:
				exitChannel <- message
			}
		}
	}()

	return exitChannel
}

func escrever(texto string, delay int) <-chan string {
	channel := make(chan string)

	// goroutine para enviar dados para o canal
	go func() {
		for i := 0; i < 10; i++ {
			channel <- texto
			time.Sleep(time.Millisecond * time.Duration(delay))
		}
		close(channel)
	}()

	return channel
}

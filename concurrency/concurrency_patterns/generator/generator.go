package main

import (
	"fmt"
	"time"
)

func main() {

	// Canal de textos
	channelText := escrever("Olá")

	// Ouvindo canal de textos
	for message := range channelText {
		fmt.Println(message)
	}

}

// Generator - Retorna um canal que em que só posso receber mensagens
func escrever(texto string) <-chan string {
	channel := make(chan string)

	// goroutine para enviar dados para o canal
	go func() {
		for i := 0; i < 10; i++ {
			channel <- texto
			time.Sleep(time.Millisecond * 500)
		}
		close(channel)
	}()

	return channel
}

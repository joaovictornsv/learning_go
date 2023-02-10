package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Channel 2"
		}
	}()

	for {
		// FORMA ONDE O ATRASO DO CANAL 2 ATRAPALHA OS PRINTS DO CANAL 1
		// messageChannel1 := <-channel1 // Recebe 2 mensagens por seg
		// fmt.Println(messageChannel1)

		// messageChannel2 := <-channel2 // Recebe 1 mensagem a cada 2 segundos
		// fmt.Println(messageChannel2)
		// Dessa forma, o programa fica dependente do delay do canal 2

		// FORMA ONDE O DELAY DO CANAL 2 NÃO ATRAPALHA
		// Pois sempre que um dos dois canais receber mensagem, ela vai ser mostrada na tela
		select {
		case messageChannel1 := <-channel1:
			fmt.Println(messageChannel1)
		case messageChannel2 := <-channel2:
			fmt.Println(messageChannel2)
		}

	}

}

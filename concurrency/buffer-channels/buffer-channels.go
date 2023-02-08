package main

import "fmt"

func main() {
	// Canal com capacidade de 2 mensagens
	// ou seja, o canal será "bloqueado", aguardando as mensagens serem recebidas
	// apenas quando ultrapassar a capacidade máxima
	channel := make(chan string, 2)

	channel <- "Message 1"
	channel <- "Message 2"
	// channel <- "Message 3" // Deadlock! Está enviando mensagem para o canal, que está no limite, passando a se tornar bloqueante. O erro reside em não ter algum trecho de código ouvindo essa terceira mensagem canal

	msg1 := <-channel
	msg2 := <-channel
	// msg3 := <-channel // Deadlock! Está esperando uma terceira mensagem, porém o canal não está recebendo mais nada
	fmt.Println(msg1)
	fmt.Println(msg2)

}

/*
Concorrência x Paralelismo
- Concorrência: dois processos revezando o tempo de processamento
- Paralelismo: dois processos sendo executados ao mesmo tempo
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutine - não vai esperar essa chamada terminar para seguir com o programa
	go escreverVariasVezes("Primeiro texto")
	escreverVariasVezes("Segundo texto")
}

func escreverVariasVezes(texto string) {
	for i := 0; i < 20; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

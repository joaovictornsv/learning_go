package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Primeiro texto") // goroutine
	escrever("Segundo texto")
}

func escrever(texto string) {
	for i := 0; i < 20; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

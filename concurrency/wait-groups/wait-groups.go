package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	// Informa ao waitGroup quantas goroutines ele deve aguardar
	waitGroup.Add(4)

	go func() {
		escrever("Goroutine 1")
		waitGroup.Done() // Informa ao waitGroup que essa goroutine foi finalizada
	}()

	go func() {
		escrever("Goroutine 2")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutine 3")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutine 4")
		waitGroup.Done()
	}()

	// Informa ao programa para aguardar as contagens das goroutines chegar a zero
	// Dessa forma, todas as goroutines executaram de forma concorrente
	waitGroup.Wait()

	// Só irá executar após a contagem de goroutines chegar em zero
	escrever("Finalizei")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

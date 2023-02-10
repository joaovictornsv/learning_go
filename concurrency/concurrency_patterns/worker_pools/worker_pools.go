package main

import "fmt"

func main() {
	tasks, results := make(chan int, 100), make(chan int, 100)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	// A partir daqui o worker (ou workers) está esperando o canal "tasks" receber algo. Quantos mais workers, mais rápido serão calculados os valores, pois os cálculos serão feitos de forma concorrente

	// Enviando todos os números para o canal "tasks"
	for i := 0; i < 40; i++ {
		fmt.Println("Sending to tasks:", i)
		tasks <- i
	}
	close(tasks) // Fechando "tasks" após enviar tudo

	// Agora será ouvido o canal "results" com os valores calculados
	for result := range results {
		fmt.Println(result)
	}
}

// <-chan = canal em que só posso receber dados
// chan<- = canal em que só posso enviar dados
func worker(tasks <-chan int, results chan<- int) {
	// Ouvindo "tasks", calculando e enviando para "results"
	for num := range tasks {
		results <- fibonacci(num)
		fmt.Println("Calculating:", num)
	}

	close(results) // Fechando "results" após todos os cálculos
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}

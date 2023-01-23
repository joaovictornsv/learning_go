package main

import "fmt"

func createIncrementer(increment int) func(num int) int {
	f := func(num int) int {
		result := num + increment
		return result
	}

	return f
}

func main() {
	oneIncrementer := createIncrementer(1)
	twoIncrementer := createIncrementer(2)
	threeIncrementer := createIncrementer(3)

	fmt.Println(oneIncrementer(1))
	fmt.Println(twoIncrementer(1))
	fmt.Println(threeIncrementer(1))
}

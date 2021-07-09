package main

import (
	"fmt"
)

func Float() {
	// FLOAT
	var f1 float32 = 3.14
	var f2 float32 = 2.1

	fmt.Println(f1 + f2)
	fmt.Println(f1 - f2)
	fmt.Println(f1 * f2)

	// Round float
	fmt.Printf("%.3f\n", (10.0 / 3))

	fmt.Println()
}

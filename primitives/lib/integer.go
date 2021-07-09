package lib

import (
	"fmt"
)

func Integer() {
	// #### INTERGER ####
	// It is not allowed to do math operations with numbers of different types
	// Ex: int + int8; float32 - int;
	// Type cast is required
	var a int = 10
	var b int8 = 3
	fmt.Println(a + int(b))
	fmt.Println(int8(a) + b)

	fmt.Println()
}

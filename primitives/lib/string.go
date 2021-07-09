package lib

import (
	"fmt"
)

func String() {
	// Important: Strings are immutable

	str := "This is a string"
	fmt.Println(str)
	fmt.Printf("str[2] = %v\n", str[2])         // Show ASCII code
	fmt.Printf("str[2] = %v\n", string(str[2])) // Show string

	bytes := []byte(str)

	fmt.Printf("[]bytes(str) = %v\n", bytes) // Array of ASCII codes

	fmt.Println()

	// Concatenation
	fmt.Println(str + str)

	fmt.Println()

	// Runes
	// Strings using double quotes ""
	// Runes using single quotes ''
	// Runes is true type alias from int32, ie both are the same thing
	var a rune = 'a' // ASCII code for a
	b := 'b'         // ASCII code for b

	fmt.Printf("a => value: %v type: %T\n", a, a)
	fmt.Printf("b => value: %v type: %T\n", b, b)

	fmt.Println()
}

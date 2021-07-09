package variables

import (
	"fmt"
	"strconv"
)

func main() {
	// Ways to declaration and initialize variables
	var number int
	number = 10

	var otherNum int = 20

	num := 10 // Go assign a type automatically

	// Format printing .Printf
	fmt.Printf("The type of %v is %T\n", number, number)
	fmt.Printf("The type of %v is %T\n", otherNum, otherNum)
	fmt.Printf("The type of %v is %T\n", num, num)

	// Convert NUMBER to STRING using strconv package
	stringNum := strconv.Itoa(number)
	fmt.Printf("The type of %v is %T\n", stringNum, stringNum)

	// Naming convention => PascalCase or camelCase

	// Visibility
	// lowercase => package level;
	// first letter uppercase => global level; variable to export

}

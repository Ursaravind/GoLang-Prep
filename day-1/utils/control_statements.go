package utils

import (
	"fmt"
)

func ControlStatements() {
	const num = 20
	if num%2 == 0 {
		fmt.Println("even number")
	} else if num%2 != 0 {
		fmt.Println("odd number")
	} else {
		fmt.Println("invalid input")
	}
}

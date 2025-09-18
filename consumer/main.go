package main

import (
	"fmt"

	"github.com/corebreaker/shared-lib/utils"
)

func main() {
	// Using math utilities
	a, b := 10, 20
	fmt.Printf("Add(%d, %d) = %d\n", a, b, utils.Add(a, b))
	fmt.Printf("Multiply(%d, %d) = %d\n", a, b, utils.Multiply(a, b))
	fmt.Printf("Max(%d, %d) = %d\n", a, b, utils.Max(a, b))

	// Using string utilities
	text := "hello world"
	fmt.Printf("Capitalize('%s') = '%s'\n", text, utils.Capitalize(text))
	fmt.Printf("Reverse('%s') = '%s'\n", text, utils.Reverse(text))
}
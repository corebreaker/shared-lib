package utils

// Add performs addition of two integers
func Add(a, b int) int {
	return a + b
}

// Multiply performs multiplication of two integers
func Multiply(a, b int) int {
	return a * b
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
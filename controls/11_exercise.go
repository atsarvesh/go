package main

import (
	"fmt"
	"math"
)

// Square root function using Newton's method
func Sqrt(x float64) float64 {
	
	z := 1.0 // z is 1

	fmt.Printf("Calculate Sqrt(%g):\n", x)

	const iterations = 10
	for i := 0; i < iterations; i++ {
		// Newton's method adjustment formula
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %2d: z = %g\n", i+1, z)
	}

	return z
}

func main() {
	
	xVal := 2.0
	myResult := Sqrt(xVal)
	stdResult := math.Sqrt(xVal)

	fmt.Printf("Results for x = %g\n", xVal)
	fmt.Printf("My Sqrt: %g\n", myResult)
	fmt.Printf("math.Sqrt: %g\n", stdResult)
	fmt.Printf("Difference: %g\n", math.Abs(myResult-stdResult))
}
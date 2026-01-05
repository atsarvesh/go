package main

import (
	"fmt"
)

// create custom type
type ErrNegativeSqrt float64

// implement Error() method to satisfy error interface

func (e ErrNegativeSqrt) Error() string {

	// convert e to float to avoid infinite loop
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0

	// newton method 
	for i := 0; i < 10; i++ {

		z -= (z*z - x) / (2 * z)
	}

	return z, nil
}

func main() {

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-12))
}



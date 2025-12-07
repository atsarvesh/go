package main

import (
	"fmt"
	"math"
)

/* 	in go, we can define methods with types

method is a function with a special receiver argument


*/

type Vertex struct {
	X, Y float64
}

// receiver appears in its own argument list between the func keyword and the method name

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// methods are functions

func Abs1(v1 Vertex) float64 {
	return math.Sqrt(v1.X*v1.X + v1.Y*v1.Y)
}

// methods on non-struct types

type MyFloat float64

func (f MyFloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	v1 := Vertex{2, 4}

	fmt.Println(v.Abs())
	fmt.Println(Abs1(v1))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs2())

}

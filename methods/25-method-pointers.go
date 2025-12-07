package main

import (
	"fmt"
	"math"
)

/*
	methods with pointer receivers

	the receiver type has the literal syntax *T for some type T (T cannot itself be a pointer such as *int

	can modify the value to which the receiver points



*/

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) { // pointer receiver 
	v.X = v.X * f
	v.Y = v.Y * f
}

// methods are written as functions
func Abs(v1 Vertex) float64 {
	return math.Sqrt(v1.X*v1.X + v1.Y*v1.Y)
}

func Scale(v1 *Vertex, f float64) { // check the difference by removing *
	v1.X = v1.X * f
	v1.Y = v1.Y * f
}

func main() {

	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	v1 := Vertex{3, 4}
	Scale(&v1, 10)
	fmt.Println(Abs(v1))
}
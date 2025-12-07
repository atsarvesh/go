package main

import (
	"fmt"
	"math"
)

/*
	functions with a pointer argument must take a pointer // &v

	methods with pointer receivers take either a value or a pointer as the receiver when they are called // v or &v

	OR

	functions that take a value argument must take a value of that specific type // v

	methods with value receivers take either a value or a pointer as the receiver when they are called // v or &v

*/

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v1 Vertex) Abs1() float64 {
	return math.Sqrt(v1.X*v1.X + v1.Y*v1.Y)
}

func AbsFunc(v1 Vertex) float64 {
	return math.Sqrt(v1.X*v1.X + v1.Y*v1.Y)
}

func main() {

	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)

	v1 := Vertex{8, 4}
	fmt.Println(v1.Abs1())
	fmt.Println(AbsFunc(v1))

	p1 := &Vertex{4, 3}
	fmt.Println(p1.Abs1())
	fmt.Println(AbsFunc(*p1))
}
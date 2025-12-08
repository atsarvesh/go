package main

import (
	"fmt"
	"math"
)

type Abser interface { // interface is a set of method signatures
	Abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}

// T now has an M() method that matches the signature required by interface I
// type T implements the interface I implicitly
func (t T) M(){
	fmt.Println(t.S)
}

func main() {
	
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// v is a Vertex (not *Vertex) and does NOT implement Abser
	/*
		The type Vertex (the value) does not implement Abser because 
		the only Abs() method available to it is defined on the pointer receiver (*Vertex)
	*/
	//a = v ~ throws compile error if uncommented

	fmt.Println(a.Abs())

	var i I = T{"Hello Interface"}
	i.M()
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

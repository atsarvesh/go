package main

import (
	"fmt"
	"math"
)

func main() {

	// type conversion ~ expression T(v) converts the value v to the type T
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// zero values ~ variables declared without an explicit initial value are given their zero value
	var i int
	var q float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, q, b, s)

	// type interface ~ variable's type is inferred from the value on the right hand side
	v := 42 
	fmt.Printf("v is of type %T\n", v)
}
/*
	demonstrates interface values

	they hold bothh concrete value and its type at run time
*/

package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

// method M() for pointer type/ receiver *T ~ implements interface I

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

// method M() for value type/ receiver F ~ implements interface I

func (f F) M() {
	
	fmt.Println(f)
}

func main() {

	var i I // variable of interface type

	i = &T{"Hello"} // assign address of T to i

	describe(i) // helper function

	i.M() // calls method through interface ~ prints hello

	i = F(math.Pi) // assigns address of value of type F to i

	describe(i) // helper function

	i.M() // returns pi value
}

func describe(i I) {
	fmt.Println("(%v, %T)\n", i, i)
}
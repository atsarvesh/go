/*
	interface value holding a nil concrete value

	when a nil concrete value (like a nil pointer) is assigned to an interface variable,
	the resulting interface value itself is non-nil because it still holds a type
*/

package main

import "fmt"

type J interface {
	M()
}

type Q struct {
	S string
}

// critical nil check

func (q *Q) M() {
	if q == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(q.S)
}

func main() {
	var j J

	var q *Q
	j = q
	describe(j)
	j.M()

	j = &Q{"hello"}
	describe(j)
	j.M()
}


// %v ~ dynamic value and %T ~ dynamic type
func describe(j J) {
	fmt.Printf("(%v, %T)\n", j, j)
}
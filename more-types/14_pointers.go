package main

import "fmt"

/*
	pointer holds the memory address of a value

	type *T is a pointer to a T value and its zero value is nil

	& operator generates a pointer to its operand

	* operator denotes the pointer's underlying value

	known as indirecting or deferencing

*/

func zeroval(xval int) {
	xval = 0
}

func zeroptr(xptr *int) {
	*xptr = 0
}

func main() {
	i, j := 42, 2701

	p := &i         // points to i

	fmt.Println(*p) // read i through the pointer

	*p = 21         // set i through the pointer

	fmt.Println(i)  // new value of i

	p = &j         // point to j

	*p = *p / 37   // divide j through the pointer

	fmt.Println(j) // new value of j

	x := 1
	fmt.Println("initial:", x)

	zeroval(x) // pass by value
	fmt.Println("zeroval:", x)

	// `&x` syntax gives the memory address of `x`, i.e. a pointer to `x`
	zeroptr(&x)
	fmt.Println("zeroptr:", x)

	// pointers can be printed too
	fmt.Println("pointer:", &x)
}

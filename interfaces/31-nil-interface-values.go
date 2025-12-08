// nil interface value

package main

import "fmt"

type I interface {
	M()
}

func main() {
	
	/*
		since i is not explicitly initialized, 
		it holds the zero value for an interface, which is nil 
		
		a nil interface has neither a concrete type nor a concrete value
	*/
	
	var i I

	describe(i)

	/*
		because i is a nil interface, 
		the runtime cannot find a concrete type or a concrete method to execute, 
		resulting in a panic
	*/

	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
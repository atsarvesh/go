package main

import "fmt"

/* 	type [n]T is an array of n values of type T

arrays can"t be resized
*/

func main() {

	var a [2]string // array of 2 strings
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// by default array is zero valued
    var x [5]int
    fmt.Println("emp:", x)

    // set a value at an index using the array[index] = value
	x[4] = 100
    fmt.Println("set:", x)
    fmt.Println("get:", x[4])

	// length o fan array
    fmt.Println("len:", len(x))

    // array in one line
	b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

	// compiler count the number of elements in an array
    b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

	// specify the index with (:) the elements in between will be zeroed
    b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)

	// multi-dimensional array ds
    var twoD [2][3]int
    for i := range 2 {
        for j := range 3 {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

    twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }
    fmt.Println("2d: ", twoD)
}
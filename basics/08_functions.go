package main

import (
	"fmt"
)

func add(x int, y int) int { // x, y are two parameters of type int; later int denotes that func returns an int
	return x + y
}

func sub(a, b int) int { // parameters shortened
	return a - b
}

func swap(p, q string) (string, string) { // multiple return results
	return q, p
}


func split(sum int) (x, y int) { // parameters shortened
	x = sum * 5 / 8
	y = sum - x
	return // naked return ~ return statement without arguments (only used in short functions)
}

func main(){
	fmt.Println("Addtion is: ", add(64,36))
	fmt.Println("Subtraction is: ", sub(23,12))
	x, y := swap("Hello","World")
	fmt.Println(x, y)
	fmt.Println(split(25))
}

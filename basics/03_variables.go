package main

import "fmt"

func main() {

    var x = "hello world" // var used to declare variable
    fmt.Println(x)

    var y, z int = 1, 2 // variables with initializers
    fmt.Println(y, z)

    var w = true
    fmt.Println(w)

    var p int
    fmt.Println(p)

    fruit := "apple" // short variable declaration (implicit type) ~ can only be used inside a function
    fmt.Println(fruit)
    
}
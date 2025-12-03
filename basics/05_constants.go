package main

import (
	"fmt"
	"math"
)

const s string = "constant" // const is uesd to declare constants (can't be declared using :=)~ string, character, boolean 

const (

	Big = 1 << 100 // create number by shifting 1 bit left 100 places or number 1 followed by 100 0s

	Small = Big >> 99
)

func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}



func main() {

    fmt.Println(s)

    const n = 500000000 // numeric constant ~ high precision values

    const d = 3e20 / n

    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))

    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))

}
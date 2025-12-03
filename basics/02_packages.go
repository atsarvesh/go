package main // program starts to run

import ( // factored import statement ~ group the imports
	"fmt"
	"math"
	"math/rand"
)

func main() {

    fmt.Println("My favourite number is", rand.Intn(10))
    fmt.Println("My favourite number is", math.Pi) // Pi ~ exported name that starts with capital letter 

}
package main

import "fmt"

/*

for loop has 3 components: (init; cond; post)
	init statement: executed before the first iteration
	condition expression: evaluated before every iteration
	post statement: executed at the end of every iteration
*/
func main() {

	// classic for loop
	for j := 0; j < 3; j++ {
        fmt.Println(j)
    }

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// init and post statements are optional
	total := 1
	for ; total < 1000; {
		total += total 
	}
	fmt.Println(total)


    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

	// do this N times
    for i := range 3 {
        fmt.Println("range", i)
    }

	// break
    for {
        fmt.Println("loop")
        break
    }

	// continue
    for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }

	// forever
	for {
	fmt.Println("hi")
	}
}
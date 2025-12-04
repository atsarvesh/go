package main

import "fmt"

/* defer ~ used to schedule a function call to be executed
   just before the surrounding function completes its execution (returns)

   working- scheduling, argument evaluation, execution

   deferred function calls are pushed onto a stack
   when a function returns, its deferred calls are executed in last-in-first-out order
*/

func main() {
    defer fmt.Println("world") // scheduled first, executes second
    fmt.Println("hello")
    defer fmt.Println("go")    // scheduled second, executes first (LIFO)
}

// Output:
// hello
// go
// world

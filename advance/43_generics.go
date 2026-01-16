/*
	generics allow you to write functions and data structures that can operate on different types while providing compile-time type safety.
*/

package main

import (
	"cmp"
	"fmt"
)

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func SlicesIndex[S ~[]E, E comparable](s S, v E) int { // ~ indicates a type constraint that S must be a slice of elements of type E
    for i := range s {
        if v == s[i] {
            return i
        }
    }
    return -1
}

type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val  T
}

func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func main() {
	
	// using the generic Max function
	fmt.Println("Max of 3 and 5:", Max(3, 5))
	fmt.Println("Max of 2.7 and 3.4:", Max(2.7, 3.4))
	fmt.Println("Max of 'apple' and 'banana':", Max("apple", "banana"))


    var s = []string{"foo", "bar", "zoo"}

    fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

    _ = SlicesIndex[[]string, string](s, "zoo")

    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.AllElements())
}
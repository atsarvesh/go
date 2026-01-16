/*

 */

package main

import (
	"fmt"
	"iter"
	"slices"
)

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

func (lst *List[T]) AllElements() iter.Seq[T] {
	return func (yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}	
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {

	lst := List[int]{}
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)

	fmt.Println("elements in the list:")

	for v := range lst.AllElements() {
		fmt.Println(v)
	}

	all := slices.Collect(lst.AllElements())
	fmt.Println("all elements collected into a slice:", all)

	for c := range genFib() {
		if c > 100 {
			break
		}
		fmt.Println(c)
	}
}
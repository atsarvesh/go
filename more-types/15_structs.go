package main

import "fmt"

/* 	struct is a collection of fields ~ group data to form records

struct fields are accessed using (.) dot
*/

// struct type

type Vertex struct {
	X int
	Y int
}

type person struct {
    name string
    age  int
}

// constructor function ~ useful in garbage collection

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 18
	return &p
}

// struct literals ~ newly allocated struct value

var (
	v1 = Vertex{2,5} // has type vertex
	v2 = Vertex{X: 6} // Y: 0 is implicit
	v3 = Vertex{} // X: 0 and Y: 0
	x = &Vertex{1,2} // has type
)

func main() {

	fmt.Println(Vertex{1, 2})

	v := Vertex{7, 5}
	v.X = 4
	fmt.Println(v.X)

	// struct fields can be accessed through struct pointers (*p).X or just p.X
	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, v2, v3, x)

	// creates a new struct
	fmt.Println(person{"Bob", 20})

	// name the fields when initializing a struct
	fmt.Println(person{name: "Alice", age: 30})

	// omitted fields will be 0
	fmt.Println(person{name: "Fred"})

	// `&` prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	// access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

	// anonymous structs ~ used in one-off values
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}

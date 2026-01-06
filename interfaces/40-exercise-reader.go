package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct {}

// Read populates given byte slice with 'A' & returns length
func (r MyReader) Read(b []byte) (int, error){

	// fill every available slot in slice with 'A'
	for i := range b {
		b[i] = 'A'
	}

	// return no. of bytes populated & no error
	return len(b), nil
}

func main() {

	reader.Validate(MyReader{})
}
package main

import "fmt"

// Qux is a struct with contents.
type Qux struct {
	// Contents is the slice which contains the contents of the Qux struct.
	Contents []string
}

// PrintContents prints the contents of the Qux struct
func (q *Qux) PrintContents() {
	fmt.Printf("%v", q.Contents)
}

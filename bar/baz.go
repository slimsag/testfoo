package bar

import "fmt"

// Baz is a struct with contents.
type Baz struct {
	// Contents is the slice which contains the contents of the Baz struct.
	Contents []string
}

// PrintContents prints the contents of the Baz struct
func (b *Baz) PrintContents() {
	fmt.Printf("%v", b.Contents)
}

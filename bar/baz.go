package bar

import "fmt"

type Baz struct {
	Contents []string
}

func (b *Baz) PrintContents() {
	fmt.Printf("%v", b.Contents)
}

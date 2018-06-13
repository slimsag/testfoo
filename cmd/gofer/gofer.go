package main // import "code.slimsag.com/the/repo/cmd/gofer"

import "code.slimsag.com/the/repo/bar"

func main() {
	b := bar.Baz{
		Contents: []string{
			"Ham",
			"Eggs",
		},
	}

	b.PrintContents()

	q := Qux{
		Contents: []string{
			"Spam",
			"Green",
		},
	}

	q.PrintContents()
}

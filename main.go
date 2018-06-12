package main // import "whoami.com/FOO"

import "whoami.com/foo/bar"

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

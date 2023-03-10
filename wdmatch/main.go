package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args

	if len(a) == 3 {
		b := os.Args[1]
		c := os.Args[2]
		d := []rune{}

		pos := 0

		for i := 0; i < len(b); i++ {
			for k := pos; k < len(c); k++ {
				if b[i] == c[k] {
					d = append(d, rune(c[k]))
					pos = k
					break
				}
			}
		}
		if string(d) == b {
			for n := 0; n < len(d); n++ {
				z01.PrintRune(d[n])
			}
			z01.PrintRune('\n')
		}

	}
}

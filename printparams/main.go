package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	c := os.Args

	for i := 1; i < len(c); i++ {
		for k := 0; k < len(c[i]); k++ {
			z01.PrintRune(rune(c[i][k]))
		}
		z01.PrintRune('\n')
	}
}

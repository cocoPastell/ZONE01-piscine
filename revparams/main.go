package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args
	compteur := 0

	for range a {
		compteur++
	}

	for i := compteur - 1; i >= 1; i-- {
		for k := 0; k < len(a[i]); k++ {
			z01.PrintRune(rune(a[i][k]))
		}
		z01.PrintRune('\n')
	}
}

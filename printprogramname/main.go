package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	c := []rune(os.Args[0])

	for i := 2; i < len(c); i++ {
		z01.PrintRune(c[i])
	}
	z01.PrintRune('\n')
}

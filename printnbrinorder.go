package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	c := []int{}

	if n == 0 {
		z01.PrintRune('0')
	}

	for n > 0 {

		c = append(c, n%10)
		n = n / 10
	}

	for i := 0; i < len(c); i++ {
		for k := i + 1; k < len(c); k++ {
			if c[i] > c[k] {
				a := c[i]
				c[i] = c[k]
				c[k] = a
			}
		}
	}
	for i := 0; i < len(c); i++ {
		z01.PrintRune(rune(c[i] + 48))
	}
}

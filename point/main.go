package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func setNumber(n int) {
	if n < 10 {
		for i, k := 0, '0'; i <= n; i, k = i+1, k+1 {
			if i == n {
				z01.PrintRune(k)
			}
		}
	} else {
		nbr := n / 10
		mod := n % 10
		setNumber(nbr)
		setNumber(mod)
	}
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	setNumber(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	setNumber(points.y)
	z01.PrintRune(10)
}

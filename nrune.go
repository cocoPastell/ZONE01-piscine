package piscine

func NRune(s string, n int) rune {
	a := []rune(s)
	c := 0

	for range s {
		c++
	}
	if n < 1 || n > c {
		return 0
	}
	return (a[n-1])
}

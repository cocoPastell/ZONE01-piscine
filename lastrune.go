package piscine

func LastRune(s string) rune {
	a := []rune(s)
	c := 0
	for range s {
		c++
	}
	return (a[c-1])
}

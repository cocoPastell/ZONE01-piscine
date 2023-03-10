package piscine

func ToUpper(s string) string {
	c := []rune(s)
	for i := 0; i < len(s); i++ {
		if c[i] >= 'a' && c[i] <= 'z' {
			c[i] -= 'a' - 'A'
		}
	}
	return string(c)
}

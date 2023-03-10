package piscine

func ToLower(s string) string {
	c := []rune(s)

	for i := 0; i < len(s); i++ {
		if c[i] >= 'A' && c[i] <= 'Z' {
			c[i] += 'a' - 'A'
		}
	}
	return string(c)
}

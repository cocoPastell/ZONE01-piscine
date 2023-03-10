package piscine

func Rot14(s string) string {
	c := []rune(s)

	for i := 0; i < len(s); i++ {
		if (c[i] >= 'A' && c[i] <= 'L') || (c[i] >= 'a' && c[i] <= 'l') {
			c[i] = c[i] + 14
		} else if (c[i] >= 'M' && c[i] <= 'Z') || (c[i] >= 'm' && c[i] <= 'z') {
			c[i] = c[i] - 12
		}
	}
	return string(c)
}

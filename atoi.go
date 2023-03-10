package piscine

func Atoi(s string) int {
	pos := 0
	c := 0

	if s[0] == '-' || s[0] == '+' {
		pos = pos + 1
	}

	for i := pos; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		if s[0] == '-' {
			c = c*10 - int(s[i]-'0')
		} else {
			c = c*10 + int(s[i]-'0')
		}
	}
	return c
}

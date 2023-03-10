package piscine

func TrimAtoi(s string) int {
	c := 0
	var a string
	var b string

	if len(s) == 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' || s[i] == '-' {
			a = a + string(s[i])
		}
	}
	if len(a) != 0 {
		if a[0] == '-' {
			for i := 1; i < len(a); i++ {
				c = c*10 - int(a[i]-'0')
			}
			return c
		} else {
			for i := 0; i < len(a); i++ {
				if a[i] >= '0' && a[i] <= '9' {
					b = b + string(a[i])
				}
			}
		}
	}

	for i := 0; i < len(b); i++ {
		c = c*10 + int(b[i]-'0')
	}
	return c
}

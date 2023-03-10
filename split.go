package piscine

func Split(s, sep string) []string {
	c := []string{}
	b := len(sep)
	x := ""

	for i := 0; i < len(s); i++ {
		if i < len(s)-b {
			if s[i:i+b] != sep {
				x = x + string(s[i])
			}
			if s[i:i+b] == sep {
				c = append(c, string(x))
				x = ""
				i = i + b - 1
			}
		} else {
			x = x + string(s[i])
			if s[i] == s[len(s)-1] {
				c = append(c, x)
			}

		}
	}

	return c
}

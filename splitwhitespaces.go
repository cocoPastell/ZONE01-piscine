package piscine

func SplitWhiteSpaces(s string) []string {
	/* 	c := []string{}
	   	x := ""

	   	for i := 0; i < len(s); i++ {
	   		if string(s[i]) != " " && string(s[i]) != "\t" && string(s[i]) != "\n" {
	   			x = x + string(s[i])
	   		}
	   		if string(s[i]) == " " || string(s[i]) == string(s[len(s)-1]) || string(s[i]) == "\t" || string(s[i]) == "\n" {
	   			if x != "" {
	   				c = append(c, string(x))
	   				x = ""
	   			}
	   		}

	   	}
	   	return c */

	if len(s) > 0 {
		a := []rune(s)
		v := []string{}
		n := ""

		for i := 0; i < len(s); i++ {
			if (a[i] != ' ' && a[i] != '\t' && a[i] != '\n') && i < len(s)-1 {
				n = n + string(a[i])
			} else {
				if n != "" {
					v = append(v, n)
					n = ""
				}
			}
		}
		p := len(s) - 1
		f := string(s[p])
		e := len(v) - 1
		v[e] = v[e] + f
		return v
	}
	return nil
}

package piscine

import "os"

func ItoA(n int) string {
	s := ""
	x := false

	if n < 0 {
		n = n * -1
		x = true
	}
	for n > 0 {
		s = string(n%10+48) + s
		n = n / 10
	}
	if x == true {
		s = "-" + s
	}
	return s
}

func AtoI(s string) int {
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

func main() {
	a := os.Args

	if len(a) == 4 {

		b := AtoI(os.Args[1])
		c := os.Args[2]
		d := AtoI(os.Args[3])

		if c == "%" && d == 0 {
			os.Stdout.WriteString("No modulo by 0")
			// os.Stdout.WriteString("\n")
		}
		if c == "/" && d == 0 {
			os.Stdout.WriteString("No division by 0")
			// os.Stdout.WriteString("\n")
		}
		if (b > -9223372036854775807 && b < 9223372036854775807) && (d > -9223372036854775807 && d < 9223372036854775807) {
			e := 0
			if c == "+" {
				e = b + d
			}
			if c == "-" {
				e = b - d
			}
			if c == "*" {
				e = b * d
			}
			if c == "/" && d != 0 {
				e = b / d
			}
			if c == "%" && d != 0 {
				e = b % d
			}
			os.Stdout.WriteString(ItoA(e))
			os.Stdout.WriteString("\n")

		}
	}
}

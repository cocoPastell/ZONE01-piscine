package piscine

func Itoa(n int) string {
	c := ""
	neg := false

	if n < 0 {
		n = n * -1
		neg = true

	}

	for n > 0 {
		c = string(n%10+48) + c
		n = n / 10
	}
	if neg == true {
		c = "-" + c
	}

	return c
}

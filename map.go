package piscine

func Map(f func(int) bool, a []int) []bool {
	var c []bool
	for _, i := range a {
		c = append(c, f(i))
	}
	return c
}

package piscine

func AppendRange(min, max int) []int {
	c := []int{}

	if min >= max {
		return []int(nil)
	}

	for i := min; i < max; i++ {
		c = append(c, i)
	}
	return c
}

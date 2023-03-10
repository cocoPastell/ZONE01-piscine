package piscine

func MakeRange(min, max int) []int {
	if min < max {
		c := make([]int, max-min)
		for i := 0; i < max-min; i++ {
			c[i] = min + i
		}
		return c
	}
	return []int(nil)
}

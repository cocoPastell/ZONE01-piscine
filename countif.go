package piscine

func CountIf(f func(string) bool, tab []string) int {
	c := 0
	for _, i := range tab {
		if f(i) {
			c++
		}
	}
	return c
}

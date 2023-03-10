package piscine

func Unmatch(a []int) int {
	for _, c := range a {
		compteur := 0
		for _, k := range a {
			if c == k {
				compteur++
			}
		}
		if compteur == 1 || compteur%2 == 1 {
			return c
		}
	}
	return -1
}

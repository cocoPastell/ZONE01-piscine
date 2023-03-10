package piscine

func Max(a []int) int {
	for i := 0; i < len(a); i++ {
		for k := i + 1; k < len(a); k++ {
			if a[i] < a[k] {
				c := a[i]
				a[i] = a[k]
				a[k] = c
			}
		}
	}
	return a[0]
}

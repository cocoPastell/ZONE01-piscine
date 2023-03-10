package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 1 {
		return 1
	} else {
		for i := 1; i <= nb; i++ {
			if nb == i*i {
				return i
			}
		}
	}
	return 0
}

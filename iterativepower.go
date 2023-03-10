package piscine

func IterativePower(nb int, power int) int {
	k := 1
	if power < 0 {
		return 0
	} else {
		for i := power; i > 0; i-- {
			k = k * nb
		}
	}
	return k
}

func ActiveBits(n int) int {
	compteur := 0

	for n != 0 {
		if n%2 == 1 {
			compteur++
		}
		n = n / 2

	}

	return compteur
}

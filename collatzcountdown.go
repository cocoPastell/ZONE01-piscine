package piscine

func CollatzCountdown(start int) int {
	compteur := 0

	if start <= 0 {
		return -1
	} else {
		for start != 1 {
			if start%2 == 0 {
				start = start / 2
				compteur++

			} else {
				start = start*3 + 1
				compteur++
			}
		}
	}
	return compteur
}

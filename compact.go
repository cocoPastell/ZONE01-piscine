package piscine

func Compact(ptr *[]string) int {
	var k []string
	compteur := 0

	for _, c := range *ptr {
		if c != "" {
			k = append(k, c)
			compteur++
		}
	}
	*ptr = k
	return compteur
}

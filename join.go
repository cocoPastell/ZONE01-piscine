package piscine

func Join(strs []string, sep string) string {
	var c string

	for i, k := range strs {
		c = c + k
		if i == len(strs)-1 {
			break
		}
		c = c + sep
	}
	return c
}

package piscine

func Index(s string, toFind string) int {
	a := len(s)
	b := len(toFind)

	for i := 0; i < a-b; i++ {
		if s[i:i+b] == toFind {
			return i
		}
	}
	return -1
}

package piscine

func ListSize(l *List) int {
	compteur := 0

	for l.Head != nil {
		compteur++
		l.Head = l.Head.Next
	}
	return compteur
}

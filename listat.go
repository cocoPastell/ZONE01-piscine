package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	compteur := 0
	for l != nil {
		if compteur == pos {
			return l
		}
		compteur++
		l = l.Next
	}
	return nil
}

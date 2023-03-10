package piscine

func ListReverse(l *List) {
	c := l.Head
	nx := l.Head
	now := l.Head
	last := l.Head
	last = nil

	for now != nil {
		nx = now.Next
		now.Next = last
		last = now
		now = nx
	}
	l.Head = l.Tail
	l.Tail = c
}

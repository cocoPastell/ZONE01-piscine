package piscine

func ForEach(f func(int), a []int) {
	for _, c := range a {
		f(c)
	}
}

package piscine

func UltimateDivMod(a *int, b *int) {
	x := *a
	y := *b
	*a = x / y
	*b = x % y
}

package piscine

func IterativeFactorial(nb int) int {
	arg := nb
	if nb < 0 || nb > 32767 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		for i := nb - 1; i > 0; i = i - 1 {
			arg = arg * i
		}
	}
	return arg
}

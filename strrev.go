package piscine

func StrRev(s string) string {
	var f string
	for i := len(s) - 1; i >= 0; i-- {
		f = f + string(s[i])
	}
	return f
}

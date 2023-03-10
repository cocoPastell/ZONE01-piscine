package piscine

func BasicAtoi(s string) int {
	k := 0

	for i := 0; i < len(s); i++ {
		if int(s[i]-'0') < 0 || int(s[i]-'0') > 9 {
			return 0
		} else {
			k = k*10 + int(s[i]-'0')
		}
	}
	return k
}

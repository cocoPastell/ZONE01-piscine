package piscine

func ReverseBits(oct byte) byte {
	c := make([]byte, 8)
	pos := 0

	for oct != 0 {
		c[pos] = oct % 2
		oct = oct / 2
		pos++
	}

	var x byte = 0
	for _, i := range c {
		x = x*2 + i
	}
	return x
}

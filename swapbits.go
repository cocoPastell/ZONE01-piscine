package piscine

func SwapBits(octet byte) byte {
	c := make([]byte, 8)

	pos := 0

	for octet != 0 {
		c[pos] = octet % 2
		octet = octet / 2
		pos++
	}

	var x byte = 0

	for i := 3; i < len(c); i++ {
		x = x + c[i]
	}
	for i := 0; i < 4; i++ {
		x = x + c[i]
	}

	return x
}

/*
func main() {
	var w byte = 100
	var r byte = SwapBits(w)
	fmt.Printf("%08b\n", r)
} */

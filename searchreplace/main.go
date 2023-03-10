package piscine

import "os"

func main() {
	a := os.Args
	b := []rune(a[1])
	c := []rune(a[2])
	d := []rune(a[3])
	for i := 0; i < len(b); i++ {
		for k := 0; k < len(c); k++ {
			for n := 0; n < len(d); n++ {
				if b[i] == c[k] {
					b[i] = d[n]
				}
			}
		}
	}
	os.Stdout.WriteString(string(b))
}

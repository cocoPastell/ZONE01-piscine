package piscine

import "github.com/01-edu/z01"

func PrintComb2() {

	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for x := '0'; x <= '9'; x++ {
					if i <= k && j < x && j <= k {
				
				if i == '9' && j == '9' && k == '9' && k == '9'{
						break
					}	
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(x)

					if i == '9' && j == '8' && k == '9' && k == '9'{
						break
					}

					z01.PrintRune(',')
					z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

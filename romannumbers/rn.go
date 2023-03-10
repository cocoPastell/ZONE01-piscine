package main

import "os"

func Atoi(s string)int{
	pos := 0
	c := 0

	if s[0] == '-' || s[0] == '+' {
		pos = pos + 1
	}

	for i := pos; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		if s[0] == '-' {
			c = c*10 - int(s[i]-'0')
		} else {
			c = c*10 + int(s[i]-'0')
		}
	}
	return c
}

func main() {
 
c := Atoi(os.Args[1])
var x string = ""
var y string = ""
var n int = 0 
len := 0

if c > 0 && c < 4000 {

	for c != 0{
	n = c%10
	c = c/10
	len++
    
	if len == 1 {
			if n == 1 {
			y = "I"	
			x = "I"
			}else if n == 2 {
				y = "I+I"
				x = "II"
			}else if n == 3 {
				y = "I+I+I"
				x = "III"
			}else if n == 4 {
				y = "(V-I)"
				x = "IV"
			}else if n == 5 {
				y = "V"
				x = "V"
			}else if n == 6 {
				y = "V+I"
				x = "VI"
			}else if n == 7 {
				y = "V+I+I"
				x = "VII"
			}else if n == 8 {
				y = "V+I+I+I"
				x = "VIII"
			}else if n == 9 {
				y = "(X-I)"
				x = "IX"
			}
		}
	if len == 2 {
			if n == 1 {
			y = ("X+") + y	
			x = ("X") + x
			}else if n == 2 {
				y = ("X+X+") + y
				x = ("XX") + x
			}else if n == 3 {
				y = ("X+X+X+") + y
				x = ("XXX") + x
			}else if n == 4 {
				y = ("(L-X)+") + y
				x = ("XL") + x
			}else if n == 5 {
				y = ("L+") + y
				x = ("L") + x
			}else if n == 6 {
				y = ("L+X+") + y
				x = ("LX") + x
			}else if n == 7 {
				y = ("L+X+X+") + y
				x = ("LXX") + x
			}else if n == 8 {
				y = ("L+X+X+X+") + y
				x = ("LXXX") + x
			}else if n == 9 {
				y = ("(C-X)+") + y
				x = ("XC") + x
			}
		}
	if len == 3 {
			if n == 1 {
			y = ("C+") + y
			x = ("C") + x
			}else if n == 2 {
				y = ("C+C+") + y
				x = ("CC") + x
			}else if n == 3 {
				y = ("C+C+C+") + y
				x = ("CCC") + x
			}else if n == 4 {
				y = ("(D-C)+") + y
				x = ("CD") + x
			}else if n == 5 {
				y = ("D+") + y
				x = ("D") + x
			}else if n == 6 {
				y = ("D+C+") + y
				x = ("DC") + x
			}else if n == 7 {
				y = ("D+C+C+") + y
				x = ("DCC") + x
			}else if n == 8 {
				y = ("D+C+C+C+") + y
				x = ("DCCC") + x
			}else if n == 9 {
				y = ("(M-C)+") + y
				x = ("CM") + x
			}
		}
	if len == 4 {
			if n == 1 {
			y = ("M+") + y
			x = ("M") + x
			}else if n == 2 {
				y = ("M+M+") + y
				x = ("MM") + x
			}else if n == 3 {
				y = ("M+M+M+") + y
				x = ("MMM") + x
			}
		}
		}
	os.Stdout.WriteString(y)
	os.Stdout.WriteString("\n")	
	os.Stdout.WriteString(x)
	os.Stdout.WriteString("\n")	
	
		
	} else {
	os.Stdout.WriteString("ERROR: cannot convert to roman digit")
	os.Stdout.WriteString("\n")	
}
}  
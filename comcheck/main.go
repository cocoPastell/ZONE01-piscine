package main

import (
	"fmt"
	"os"
)

func main() {
	c := os.Args

	for i := 1; i < len(c); i++ {
		if c[i] == "01" || c[i] == "galaxy" || c[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}

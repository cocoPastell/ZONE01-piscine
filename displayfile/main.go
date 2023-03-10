package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 2 {

		file, err := ioutil.ReadFile("quest8.txt")
		if err != nil {
			fmt.Printf("Error : %v", err.Error())
		}
		fmt.Println(string(file))

	} else if len(os.Args) < 2 {
		fmt.Println("File name missing")
	} else {
		fmt.Println("Too many arguments")
	}
}

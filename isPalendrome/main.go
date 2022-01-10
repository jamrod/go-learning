package main

import (
	"fmt"
	"os"

	"../revstring"
)

func main() {
	inputString := os.Args[1]
	reversed := revstring.ReverseString(inputString)
	if reversed == inputString {
		fmt.Println("Palendrome")
	} else {
		fmt.Println("Not a Palendrome")
	}
}

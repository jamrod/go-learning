package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	instring := os.Args[1]
	fmt.Println(instring)
	out := ""
	chars := strings.Split(instring, "")
	for i := len(chars) - 1; i >= 0; i-- {
		char := chars[i]
		out += char
	}
	fmt.Println(out)
}

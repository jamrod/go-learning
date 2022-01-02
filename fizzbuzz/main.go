package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		line := ""
		if i%3 == 0 {
			line = "Fizz"
		}
		if i%5 == 0 {
			line += "Buzz"
		}
		if len(line) < 1 {
			line = strconv.Itoa(i)
		}
		fmt.Println(line)
	}
}

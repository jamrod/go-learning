package revstring

import (
	"strings"
)

// func main() {
// 	fmt.Println(ReverseString(os.Args[1]))
// }

func ReverseString(str string) string {
	instring := str
	out := ""
	chars := strings.Split(instring, "")
	for i := len(chars) - 1; i >= 0; i-- {
		char := chars[i]
		out += char
	}
	return out
}

package main

import (
	"z01"
)

func PrintStr(s string) {
	for _, letter := range s {
		z01.PrintRune(letter)
	}

}
func StrRev(s string) string {
	var str string
	arr := make([]byte, len(s))
	var lenght int = len(s) - 1
	arr = []byte(s)
	for i := 0; i < len(s); i++ {
		arr[i] = s[lenght]
		lenght--
	}
	str = string(arr)
	return str
}

func main() {

}

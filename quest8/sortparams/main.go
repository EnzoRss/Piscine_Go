package main

import (
	"os"
	"z01"
)

func SortWordArr(a []string) []string {
	str := make([]string, len(a))
	var min rune
	var stock int
	var str1 string
	var str2 string
	var verif bool
	r := make([]rune, len(a)+1)
	r1 := make([]rune, len(a)+1)
	for i := 1; i < len(a); i++ {
		str1 = a[i]
		r = []rune(str1)
		min = r[0]
		for j := i + 1; j < len(a); j++ {
			str2 = a[j]
			r1 = []rune(str2)
			if min > r1[0] {
				min = r1[0]
				stock = j
				verif = true
			}
		}
		if verif {
			str1 = a[stock]
			str2 = a[i]
			a[i] = str1
			a[stock] = str2
		}
		verif = false
	}
	str = a
	return str
}
func PrintStr(s string) {
	for _, letter := range s {
		z01.PrintRune(letter)
	}

}

func main() {
	var str []string = SortWordArr(os.Args)

	for i := 1; i < len(os.Args); i++ {
		PrintStr(str[i])
		PrintStr("\n")
	}
}

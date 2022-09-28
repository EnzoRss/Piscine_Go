package main

import (
	"fmt"
	"os"
)

func main() {
	var str1 string = os.Args[len(os.Args)-1]
	var str2 []string
	if len(os.Args) < 2 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Println("--insert")
		fmt.Println("-i")
		fmt.Println("         This flag inserts the string into the string passed as argument.")
		fmt.Println("--order")
		fmt.Println("-o")
		fmt.Println("	  This flag will behave like a boolean, if it is called it will order the argument.")
		for i := 1; i < len(os.Args); i++ {
			str := Split(os.Args[i], "=")
			if str[0] == "--insert" || str[0] == "-i" {
				str1 = str1 + str[1]
			} else if os.Args[i] == "--order" || os.Args[i] == "-o" {
				str2 = SortWordArr(StringToArr(str1))
			}

		}
		str2 = append(str2, str1)
		for j := 0; j < len(str2)-1; j++ {
			fmt.Print(str2[j])
		}
	}

}

func Split(s, sep string) []string {

	var str string
	var str1 []string
	var str2 string
	var test bool
	var j int = 0
	var pos int = 0
	for _, letter := range s {
		for i := 0; i < len(sep); i++ {
			if s[pos] == sep[j] {
				str2 = str2 + string(letter)
				test = true
				j++
				if str2 == sep {
					str1 = append(str1, str)
					str = ""
					str2 = ""
					j = 0
				}
			}
		}
		if !test {
			str = str + str2 + string(letter)
			str2 = ""
		}
		test = false
		pos++
	}
	str1 = append(str1, str)
	return str1
}

func SortWordArr(a []string) []string {
	str := make([]string, len(a))
	var min rune
	var stock int
	var str1 string
	var str2 string
	var verif bool
	r := make([]rune, len(a)+1)
	r1 := make([]rune, len(a)+1)
	for i := 0; i < len(a); i++ {
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

func StringToArr(str string) []string {
	var arr []string

	for i := 0; i < len(str); i++ {
		arr = append(arr, string(str[i]))
	}
	return arr
}

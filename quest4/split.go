package main

import "fmt"

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
					fmt.Println("K")
					j = 0
				}
			}
		}
		if !test {
			str = str + str2 + string(letter)
			str2 = ""
		}
		fmt.Println(str)
		test = false
		pos++
	}
	str1 = append(str1, str)
	return str1
}

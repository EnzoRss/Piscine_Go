package main

import "z01"

func main() {
	var i int = 0
	var letter rune = 97
	for ; i < 26; i++ {
		z01.PrintRune(letter)
		if letter >= 122 {
			z01.PrintRune('\n')
		}
		letter++
	}
}

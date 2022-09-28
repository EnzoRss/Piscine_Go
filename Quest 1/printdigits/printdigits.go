package main

import "z01"

func main() {
	var i int = 0
	var digits rune = 48
	for ; i < 10; i++ {
		z01.PrintRune(digits)
		if digits >= 57 {
			z01.PrintRune('\n')
		}
		digits++
	}
}

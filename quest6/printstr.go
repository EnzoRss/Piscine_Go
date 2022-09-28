package main

import "z01"

func PrintStr(s string) {
	for _, letter := range s {
		z01.PrintRune(letter)
	}

}

package main

import "z01"

func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {
		for _, letter := range a[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}

}

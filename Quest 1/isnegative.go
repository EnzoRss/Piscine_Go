package main

import "z01"

func isnegative(nb int) {
	if nb < 0 {
		z01.PrintRune('F')
	} else if nb >= 0 {
		z01.PrintRune('T')
	}
}

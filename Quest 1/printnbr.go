package main

import "z01"

func printnbr(n int32) {
	if n < 0 {
		n = -n
	}
	if n >= 10 {
		printnbr(n / 10)
		printnbr(n % 10)
	} else {
		z01.PrintRune(n + '0')
	}
}

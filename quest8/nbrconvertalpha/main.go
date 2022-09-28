package main

import (
	"os"
	"z01"
)

func PrintStr(s string) {
	for _, letter := range s {
		z01.PrintRune(letter)
	}
}

func main() {
	var Upper rune = 16
	var alpha rune = 48
	var stock rune = 0
	var str string

	for i := 1; i < len(os.Args); i++ {
		if os.Args[1] == "--upper" {
			alpha = Upper
		}

		for index, letter := range os.Args[i] {
			if index > 0 && os.Args[i] < "20" {
				stock = 10
			} else if index > 0 && os.Args[i] > "20" {
				stock = 20
			}
			stock += letter
		}
		if os.Args[i] != "56" {
			str = string(stock + alpha)
			PrintStr(str)
		} else {
			PrintStr(" ")
		}
		stock = 0
		str = ""
	}
	z01.PrintRune('\n')
}

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
	for i := len(os.Args) - 1; i > 0; i-- {
		PrintStr(os.Args[i])
		z01.PrintRune('\n')
	}
}

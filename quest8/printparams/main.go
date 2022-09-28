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
	for i := 1; i < len(os.Args); i++ {
		PrintStr(os.Args[i])
		z01.PrintRune('\n')
	}
}

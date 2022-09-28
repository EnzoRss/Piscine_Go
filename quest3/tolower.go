package main

import "fmt"

func ToLower(s string) string {
	var str string
	for _, letter := range s {
		if letter >= 65 && letter <= 90 {
			fmt.Println("oui ")
			letter = letter + 32
		}
		str = str + string(letter)
	}
	return str
}

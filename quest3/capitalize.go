package main

import "fmt"

func IsAlpha(s string) bool {
	counter := 0
	stock := 0
	for index, letter := range s {
		if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122) || (letter >= 48 && letter <= 57) {
			counter++
		}
		stock = index
	}
	if stock+1 == counter {
		return true
	}
	return false
}

func ToUpper(s string) string {
	var str string
	for _, letter := range s {
		if letter >= 97 && letter <= 122 || letter <= 65 && letter >= 90 {
			letter = letter - 32
		}
		str = str + string(letter)
	}
	return str
}

func IsNumeric(s string) bool {
	counter := 0
	stock := 0
	for index, letter := range s {
		if letter >= 48 && letter <= 57 {
			counter++
		}
		stock = index
	}
	if stock+1 == counter {
		return true
	}
	return false
}

func Capitalize(s string) string {
	var str1 string
	var str2 string
	var stock rune
	for _, letter := range s {
		if !IsAlpha(string(stock)) && IsAlpha(string(letter)) && !IsNumeric(string(letter)) {
			fmt.Println("n")
			str2 = ToUpper(string(letter))
			str1 += str2
		} else {
			str1 = str1 + string(letter)
		}
		stock = letter
	}
	return str1
}

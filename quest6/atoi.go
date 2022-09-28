package main

import "fmt"

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

func Atoi(s string) int {
	var result int
	var verif bool
	var signe int = 1
	var i int
	if s[0] == 43 {
		signe = 1
		i = 1
	} else if s[0] == 45 {
		signe = -1
		i = 1
	}
	for ; i < len(s); i++ {
		if s[i] > 48 && s[i] <= 57 {
			result = result*10 + int(s[i]-'0')
			verif = true
		} else if !IsNumeric(string(s[i])) {
			verif = false
			break
		}
	}
	if !verif {
		return 0
	}
	return result * signe
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

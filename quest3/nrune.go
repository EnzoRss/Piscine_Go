package main

func NRune(s string, n int) rune {
	var letter rune
	for index, temp := range s {
		if index+1 == n {
			letter = temp
		}
	}
	return letter

}

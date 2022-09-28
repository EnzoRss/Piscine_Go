package main

func LastRune(s string) rune {

	var letter rune
	for _, temp := range s {
		letter = temp
	}
	return (letter)
}

package main

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

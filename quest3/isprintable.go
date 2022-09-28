package main

func IsPrintable(s string) bool {
	counter := 0
	stock := 0
	for index, letter := range s {
		if letter >= 32 && letter <= 127 {
			counter++
		}
		stock = index
	}
	if stock+1 == counter {
		return true
	}
	return false
}

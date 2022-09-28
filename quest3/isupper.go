package main

func IsUpper(s string) bool {
	counter := 0
	stock := 0
	for index, letter := range s {
		if letter >= 65 && letter <= 90 {
			counter++
		}
		stock = index
	}
	if stock+1 == counter {
		return true
	}
	return false
}

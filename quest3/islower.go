package main

func IsLower(s string) bool {
	counter := 0
	stock := 0
	for index, letter := range s {
		if letter >= 97 && letter <= 122 {
			counter++
		}
		stock = index
	}
	if stock+1 == counter {
		return true
	}
	return false
}

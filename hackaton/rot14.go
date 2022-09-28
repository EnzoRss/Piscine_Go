package main

func Rot14(s string) string {
	var str string
	var stock rune
	for _, letter := range s {
		if letter >= 97 && letter <= 122 {
			stock = letter + 14
			if stock > 122 {
				stock = (stock - 122) + 96
			}
		} else if letter >= 65 && letter <= 90 {
			stock = letter + 14
			if stock > 90 {
				stock = (stock - 90) + 64

			}
		} else {
			stock = letter
		}
		str = str + string(stock)
		stock = 0
	}
	return str
}

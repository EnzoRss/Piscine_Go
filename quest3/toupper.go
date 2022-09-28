package main

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

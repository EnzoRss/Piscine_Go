package main

func SplitWhiteSpaces(s string) []string {
	var str []string
	var str2 string

	for _, letter := range s {
		str2 = str2 + string(letter)
		if (letter == 32) || (letter == 9) || (letter == 13) {
			str = append(str, str2)
			str2 = ""
		}
	}
	str = append(str, str2)
	return str
}

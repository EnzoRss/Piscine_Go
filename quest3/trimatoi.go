package main

func TrimAtoi(s string) int {
	result := 0
	signe := 1
	var i int

	for i < len(s) {
		if s[i] > 48 && s[i] <= 57 {
			result = result*10 + int(s[i]-'0')
		}
		if s[i] == '-' {
			signe = -1
		}
		i++
	}
	if result == 0 {
		return 0
	}
	return result * signe
}

package main

func IsAlpha(s string) bool {
	counter := 0
	stock := 0
	for index, letter := range s {
		if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122) || letter == 32 {
			counter++
		}
		stock = index
	}
	if stock+1 == counter {
		return true
	}
	return false
}

func BasicAtoi2(s string) int {
	var result int
	var verif bool
	for i := 0; i < len(s); i++ {
		if s[i] > 48 && s[i] <= 57 {
			result = result*10 + int(s[i]-'0')
			verif = true
		} else if IsAlpha(string(s[i])) {
			verif = false
			break
		}
	}
	if !verif {
		return 0
	}
	return result
}

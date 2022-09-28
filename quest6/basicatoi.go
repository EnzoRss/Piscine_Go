package main

func BasicAtoi(s string) int {
	var result int
	var verif bool
	for i := 0; i < len(s); i++ {
		if s[i] > 48 && s[i] <= 57 {
			result = result*10 + int(s[i]-'0')
			verif = true
		}
	}
	if !verif {
		return 0
	}
	return result
}

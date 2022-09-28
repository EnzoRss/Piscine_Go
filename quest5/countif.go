package main

func CountIf(f func(string) bool, tab []string) int {
	var result int
	for i := 0; i < len(tab); i++ {
		if f(tab[i]) {
			result++
		}
	}
	return result
}

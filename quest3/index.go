package main

func Index(s string, ToFind string) int {
	count := 0
	for _, letter := range s {
		for _, letters := range ToFind {
			if letters != letter {
				count++
			}
		}
	}
	if count == 0 {
		return -1
	}
	return count
}

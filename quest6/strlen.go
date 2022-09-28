package main

func StrLen(s string) int {
	var count int
	for index := range s {
		count = index
	}
	return count
}

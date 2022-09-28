package main

func Sqrt(nb int) int {
	result := 1
	for i = 0; i<nb; i= i++ {
		result = i*i
		if result == nb {
			return i
		}else if result > nb {
			return 0
		}
	}
	return 0 
}

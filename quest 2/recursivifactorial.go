package main

var result int = 1

func RecursiveFactorial(nb int) int {

	if nb == 1 {
		return 1
	}
	if nb > 1 {
		result = nb * RecursiveFactorial(nb-1)
	}
	return result
}

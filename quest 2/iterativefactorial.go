package main

var temp int = 1

func IterativeFactorial(nb int) int {
	for i := 1; i <= nb; i++ {
		temp = temp * i
	}
	return (temp)
}

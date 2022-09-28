package main

var result1 int = 1

func RecursivePower(nb int, power int) int {
	if nb == 1 {
		return 1
	}
	if power > 1 {
		result1 = (nb * result1) * RecursivePower(nb, power-1)

	}
	return result1
}

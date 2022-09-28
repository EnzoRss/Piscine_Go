package main

var results int = 1

func IterativePower(nb int, power int) int {
	for i := 1; i <= power; i++ {
		results = results * nb
	}
	return (results)
}

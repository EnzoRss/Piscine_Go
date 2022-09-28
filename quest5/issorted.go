package main

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
}

func IsSorted(f func(a, b int) int, nbr []int) bool {

	for i := 1; i < len(nbr); i++ {
		if f(nbr[i], nbr[i-1]) == 0 || f(nbr[i], nbr[i-1]) == -1 {
			return false
		}
	}
	return true
}

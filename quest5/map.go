package main

func Map(f func(int) bool, a []int) []bool {
	stock := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			stock[i] = true
		} else {
			stock[i] = false
		}
	}
	return stock
}

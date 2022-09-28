package main

func Swap(a *int, b *int) {
	var swap1 int
	var swap2 int

	swap1 = *a
	swap2 = *b
	*a = swap2
	*b = swap1
}

package main

func CollatzCountdown(start int) int {
	result := start
	count := 0
	for ; result != 1; count++ {
		if result%2 == 0 {
			result /= 2
		} else {
			result = (result * 3) + 1
		}
	}
	return count
}

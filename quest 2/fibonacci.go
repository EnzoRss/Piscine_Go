package main

func Fibonacci(index int) int {
	result := 0
	temp := 0
	suite := 1
	for i := 1; i < index; i++ {
		result = suite + temp
		temp = suite
		suite = result

	}
	if index == 1 || index == 2 {
		return 1
	}

	return result
}

package main

func MakeRange(min, max int) []int {
	var nil []int
	var answer []int

	if max > min {
		answer = make([]int, max-min)
	}
	for i := 0; i < max; i++ {
		if i >= min {
			answer[i-min] = i
		}
	}
	if min >= max {
		return nil
	}

	return answer
}

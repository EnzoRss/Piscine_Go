package main

func AppendRange(min, max int) []int {

	var result []int
	var nil []int
	for i := 0; i < max; i++ {
		if i >= min {
			result = append(result, i)
		}
	}
	if min >= max {
		return nil
	}
	return result
}

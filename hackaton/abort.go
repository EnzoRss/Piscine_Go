package main

func Abort(a, b, c, d, e int) int {
	inf := 0
	sup := 0
	var arr []int = []int{a, b, c, d, e}
	for i := 0; i < len(arr); i++ {
		for _, nbr := range arr {
			if nbr < arr[i] {
				inf++
			} else if nbr > arr[i] {
				sup++
			}
		}
		if inf == 2 && sup == 2 {
			return arr[i]
		}
		inf = 0
		sup = 0
	}

	return 0
}

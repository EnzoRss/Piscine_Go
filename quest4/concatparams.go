package main

func ConcatParams(args []string) string {
	var str string
	for i := 0; i < len(args); i++ {
		for _, letter := range args[i] {
			str = str + string(letter)
		}
		str = str + "\n"
	}
	return str
}

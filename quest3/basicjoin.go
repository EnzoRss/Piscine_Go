package main

func BasicJoin(elems []string) string {
	var str string
	var i int = 0
	for i < len(elems) {
		str = str + elems[i]
		i++
	}
	return str
}

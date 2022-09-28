package main

func Compact(ptr *[]string) int {

	str := *ptr
	for i := range *ptr {
		if str[i] == "" {
			append(str[:i], str[i+1:]...)
		}
	}

	*ptr = str
}

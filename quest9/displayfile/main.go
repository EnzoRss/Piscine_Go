package main

import (
	"fmt"
	"os"
)

func main() {
	lenght := len(os.Args)
	data, _ := os.ReadFile(os.Args[1])

	if lenght < 2 {
		fmt.Println("File name missing")
	} else if lenght == 2 {
		fmt.Println(string(data))
	} else if lenght > 2 {
		fmt.Println("Too many arguments")
	}
}

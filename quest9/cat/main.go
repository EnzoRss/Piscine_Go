package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	lentght := len(os.Args)
	if lentght == 1 {
		scan, _ := fmt.Scan(os.Stdin)
		fmt.Println(scan)
	} else {
		for i := 1; i < lentght; i++ {
			data, _ := os.ReadFile(os.Args[i])
			if !strings.Contains(os.Args[i], ".") {
				fmt.Println("ERROR: open abc: no such file or directory")
			}
			fmt.Println(string(data))
		}
	}
}

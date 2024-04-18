package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	i2 := len(argument) - 1

	for i2 > 0 {
		os := os.Args[i2]
		ros := []rune(os)
		for i := range ros {
			z01.PrintRune(ros[i])
			if len(ros) == i+1 {
				z01.PrintRune('\n')
				i2--
				i = 0
			}
		}
		if i2 == 0 {
			break
		}
	}
}

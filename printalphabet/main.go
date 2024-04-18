package main

import "github.com/01-edu/z01"

func main() {
	for char := 97; char <= 122; char++ {
		z01.PrintRune(rune(char))
	}
	z01.PrintRune('\n')
}

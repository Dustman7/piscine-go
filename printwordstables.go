package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, w := range a {
		SIUUU := w
		for _, l := range SIUUU {
			z01.PrintRune(l)
		}
		z01.PrintRune('\n')
	}
}

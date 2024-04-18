package main

import "github.com/01-edu/z01"

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return // if variable true/false*, it's gonna stop\*
	} else {
		z01.PrintRune('/')
		if x > 1 {
			for i := 0; i < x-2; i++ {
				z01.PrintRune('*') // for is true then gonna print until i reach x-2 = 0\*
			}
			z01.PrintRune('\\')
		}

		if y > 1 {
			z01.PrintRune('\n')
			for j := 0; j < y-2; j++ {
				z01.PrintRune('*') // same as P1 with some space for getting to the goal\*
				for j := 0; j < x-2; j++ {
					z01.PrintRune(' ')
				}
				if x > 1 {
					z01.PrintRune('*')
				}
				z01.PrintRune('\n')

			}
			z01.PrintRune('\\')
			if x > 1 {
				for i := 0; i < x-2; i++ { // same as P1 of code\*
					z01.PrintRune('*')
				}
				z01.PrintRune('/')
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	QuadB(5, 3)
}

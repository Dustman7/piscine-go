package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

/*func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}*/
func PrintNum(num int) {
	i := '0'
	if num == 0 {
		z01.PrintRune('0')
		return
	}
	for j := 1; j <= num%10; j++ {
		i++
	}
	for j := -1; j >= num%10; j-- {
		i++
	}
	if num/10 != 0 {
		PrintNum(num / 10)
	}
	z01.PrintRune(i)
	return
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	PrintNum(n)
}

func main() {
	point := point{}

	point.x = 42
	point.y = 21

	// setPoint(points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	PrintNbr(point.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	PrintNbr(point.y)
	z01.PrintRune('\n')
	// z01.PrintRune("x = %d, y = %d\n", points.x, points.y)
}

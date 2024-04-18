package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var a []int
	j := 0

	if n > 0 {
		for b := n; b > 0; b = b / 10 {
			a = append(a, b%10)
		}

		for range a {
			j++
		}

		/*for i := 0; i <= len(a)-1; i++ {
			z01.PrintRune(rune(a[i] + '0'))
		}*/

		for c := 0; c < j; c++ {
			for d := 0; d < j; d++ {
				if a[c] < a[d] {
					a[d], a[c] = a[c], a[d]
				}
			}
		}

		for _, n := range a {
			z01.PrintRune(rune(n + '0'))
		}
	} else {
		z01.PrintRune('0')
	}
}

/*func main() {
	PrintNbrInOrder(659)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}*/

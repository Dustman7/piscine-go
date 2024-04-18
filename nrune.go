package piscine

func NRune(s string, n int) rune {
	h := []rune(s)
	if n < 0 {
		return 0
	}
	if n > len(s) {
		return 0
	}
	curseur := 0
	for curseur < len(s) {
		curseur++
		if curseur == n {
			return rune(h[curseur-1])
		}
	}
	return 0
}

/*func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}*/

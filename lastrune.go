package piscine

func LastRune(s string) rune {
	B := s[len(s)-1]
	return rune(B)
}

/*func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}*/

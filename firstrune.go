package piscine

func FirstRune(s string) rune {
	for _, index := range s {
		return index
	}
	return 0
}

/*func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}*/

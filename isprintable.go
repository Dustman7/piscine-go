package piscine

func IsPrintable(s string) bool {
	a := []rune(s)

	for _, i := range a {
		if i < 32 || i > 126 { // 0 and 31 is correct too
			return false
		}
	}
	return true
}

/*func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))
}*/

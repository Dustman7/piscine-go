package piscine

func IsNumeric(s string) bool {
	a := []rune(s)

	for _, i := range a {
		if i < 48 || i > 57 {
			return false
		}
	}
	return true
}

/*func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}*/

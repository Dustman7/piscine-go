package piscine

func IsUpper(s string) bool {
	a := []rune(s)

	for _, i := range a {
		if i < 65 || i > 90 {
			return false
		}
	}
	return true
}

/*func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
}*/

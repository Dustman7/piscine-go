package piscine

func IsLower(s string) bool {
	a := []rune(s)

	for _, i := range a {
		if i < 97 || i > 122 {
			return false
		}
	}
	return true
}

/*func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
}*/

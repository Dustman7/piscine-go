package piscine

func IsAlpha(s string) bool {
	a := []rune(s)

	for _, i := range a {
		if (i < 65 || i > 90) && (i < 97 || i > 122) && (i < 48 || i > 57) {
			return false
		}
	}
	return true
}

/*func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
}*/

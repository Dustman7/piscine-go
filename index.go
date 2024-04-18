package piscine

func Index(s string, toFind string) int {
	a := []rune(s)
	b := []rune(toFind)
	c := len(a)
	d := len(b)

	for m := 0; m <= c-d; m++ {
		if toFind == s[m:m+d] {
			return m
		}
	}
	return -1
}

/*func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}*/

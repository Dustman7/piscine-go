package piscine

func AlphaCount(s string) int {
	nb := 0
	a := []rune(s)
	for _, j := range a {
		if j >= 97 && j <= 122 || j >= 65 && j <= 90 {
			nb++
		}
	}
	return nb
}

/*func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}*/

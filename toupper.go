package piscine

func ToUpper(s string) string {
	a := []rune(s)
	for index, value := range a {
		if value >= 97 && value <= 122 {
			a[index] = value - 32
		}
	}
	return string(a)
}

/*func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}*/

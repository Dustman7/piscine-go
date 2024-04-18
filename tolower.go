package piscine

func ToLower(s string) string {
	a := []rune(s)
	for index, value := range a {
		if value >= 65 && value <= 90 {
			a[index] = value + 32
		}
	}
	return string(a)
}

/*func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}*/

package piscine

func Rot14(s string) string {
	a := []rune(s)
	for index, value := range a {
		if value >= 'a' && value <= 'l' {
			a[index] = value + 14
		}
		if value >= 'm' && value <= 'z' {
			a[index] = value - 12
		}
		if value >= 'A' && value <= 'L' {
			a[index] = value + 14
		}
		if value >= 'M' && value <= 'Z' {
			a[index] = value - 12
		}

	}
	return string(a)
}

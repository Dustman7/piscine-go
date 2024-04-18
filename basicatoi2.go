package piscine

func BasicAtoi2(s string) int {
	ar := []rune(s)
	n := len(s)
	siuu := 0
	for i := 0; i < n; i++ {
		if ar[i] < '0' || ar[i] > '9' {
			return 0
		} else {
			siuu *= 10
			siuu += int(ar[i]) - '0'
		}
	}
	return siuu
}

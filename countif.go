package piscine

func CountIf(f func(string) bool, tab []string) int {
	digit := 0
	for _, s := range tab {
		if f(s) == true {
			digit++
		}
	}
	return digit
}

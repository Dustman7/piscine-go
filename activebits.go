package piscine

func ActiveBits(n int) int {
	var count int = 0
	for n > 0 {
		count++
		n = n & (n - 1) // & operande ET pour bits
	}
	return count
}

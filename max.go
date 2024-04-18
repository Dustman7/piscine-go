package piscine

func Max(a []int) int {
	result := a[0]
	for i := 0; i < len(a); i++ {
		if result < a[i] {
			result = a[i]
		}
	}
	return result
}

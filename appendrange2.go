package piscine

func AppendRange2(min, max int) []int {
	if min > max || max == 0 {
		return nil
	}
	toi := make([]int, max-min)

	for i, j := min, 0; i >= min && i < max; i, j = i+1, j+1 {
		toi[j] = i
	}
	return toi
}

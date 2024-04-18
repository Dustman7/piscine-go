package piscine

func MakeRange(min, max int) []int {
	if max < min || max == 0 {
		return nil
	}
	answer := make([]int, max-min)

	for i, j := min, 0; i >= min && i < max; i, j = i+1, j+1 {
		answer[j] = i
	}
	return answer
}

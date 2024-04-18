package piscine

func CollatzCountdown(start int) int {
	var i int
	if start <= 0 {
		return -1
	}
	for i = 0; start != 1; i++ {
		if start%2 == 0 {
			start = start / 2
		} else {
			start = (start * 3) + 1
		}
	}
	return i
}

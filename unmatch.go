package piscine

func Unmatch(a []int) int {
	for _, res := range a {
		fois := 0
		for _, chancla := range a {
			if chancla == res {
				fois++
			}
		}
		if fois == 1 || fois%2 == 1 {
			return res
		}
	}
	return -1
}

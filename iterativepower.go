package piscine

func IterativePower(nb int, power int) int {
	i := 1
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		for j := 0; j < power; j++ {
			i = i * nb
		}
		return i
	}
}

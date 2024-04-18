package piscine

func Map(f func(int) bool, a []int) []bool {
	d := make([]bool, len(a))
	for k, g := range a {
		d[k] = f(g)
	}
	return d
}

// siuuuu

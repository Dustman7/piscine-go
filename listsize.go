package piscine

func ListSize(l *List) int {
	n := l.Head
	count := 0
	for n != nil {
		count++
		n = n.Next
	}
	return count
}

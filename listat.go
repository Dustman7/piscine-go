package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	n := l
	count := 0
	for n != nil {
		if pos == count {
			return n
		}
		count++
		n = n.Next
	}
	return nil
}

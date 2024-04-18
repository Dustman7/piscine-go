package piscine

/*type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
} */

func ListReverse(l *List) {
	pos := l.Head
	av := l.Head
	av = nil
	for pos != nil {
		next := pos.Next
		pos.Next = av
		av = pos
		pos = next
	}
	n := l.Head
	l.Head = l.Tail
	l.Tail = n
}

package piscine

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
	} else {
		gg := l.Head
		for gg.Next != nil {
			gg = gg.Next
		}
		gg.Next = n
	}
}

/*func PrintList(Tail *NodeL) {
	for Tail != nil {
		Tail = Tail.Next
	}
	return
}*/

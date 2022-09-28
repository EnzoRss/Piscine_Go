package main

type NodeL struct {
	Data interface{}
	Next *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {

	iterator := l
	inc := 0
	for iterator != nil {
		if pos == inc {
			return iterator
		}
		inc++
		iterator = iterator.Next
	}
	return nil
}

package main

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {

	nvhead := l.Head

	for novohead != nil {
		if comp(nvhead.Data, ref) {

			return &nvhead.Data
		}
		nvhead = nvhead.Next
	}

	return nil
}

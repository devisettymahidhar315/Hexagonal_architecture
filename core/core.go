package core

type Node struct {
	Data string
	Next *Node
}

type Functions interface {
	Print()
	Ins_Beg()
	Ins_End()
	Len()
	Del_Beg()
	Del_End()
}

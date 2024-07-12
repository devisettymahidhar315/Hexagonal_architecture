package core

type Node struct {
	Data string
	Next *Node
}

type Functions interface {
	Print()
	Insertatstarting()
	InsertAtEnding()
	Length()
	DelAtBeg()
	DelAtEnd()
}

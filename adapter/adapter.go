package adapter

import (
	"app/core"
)

type JK struct {
	head *core.Node
}

func Init() *JK {
	return &JK{head: nil}
}

func (j *JK) Print() string {
	current := j.head
	var s string
	s = ""
	for current != nil {
		s = s + current.Data
		current = current.Next
	}
	return s
}

func (l *JK) InsertAtBeginning(data string) {
	newNode := &core.Node{Data: data, Next: l.head}
	l.head = newNode
}

func (l *JK) InsertAtEnding(data string) {
	newNode := &core.Node{Data: data, Next: nil}
	if l.head == nil {
		l.head = newNode
	} else {
		ite := l.head
		for ite.Next != nil {
			ite = ite.Next
		}
		ite.Next = newNode
	}
}

func (l *JK) Length() int {
	var len = 0
	ite := l.head
	for ite != nil {
		ite = ite.Next
		len = len + 1
	}
	return len
}

func (l *JK) DelAtBeg() string {
	if l.head == nil {
		return "no elements in the linked list"
	} else {
		l.head = l.head.Next
		return "first element is deleted successfully"
	}
}

func (l *JK) DelAtEnd() string {
	if l.head == nil {
		return "no elements in the linked list"
	} else if l.head.Next == nil {
		l.head = nil
		return "last element is deleted successfully"
	} else {
		ite := l.head
		for ite.Next.Next != nil {
			ite = ite.Next
		}
		ite.Next = nil
		return "last element is deleted successfully"
	}
}

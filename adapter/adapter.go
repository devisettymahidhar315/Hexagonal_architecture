package adapter

import (
	"app/core"
)

// Adapter represents a singly linked list with a head node.
type Adapter struct {
	head *core.Node
}

// Init initializes an empty linked list and returns an Adapter instance.
func Init() *Adapter {
	return &Adapter{head: nil}
}

// Print traverses the linked list and concatenates the Data of each node into a single string.
func (j *Adapter) Print() string {
	current := j.head
	var s string
	s = ""
	for current != nil {
		s = s + current.Data
		current = current.Next
	}
	return s
}

// InsertAtBeginning inserts a new node with the given data at the beginning of the linked list.
func (l *Adapter) InsertAtBeginning(data string) {
	newNode := &core.Node{Data: data, Next: l.head}
	l.head = newNode
}

// InsertAtEnding inserts a new node with the given data at the end of the linked list.
func (l *Adapter) InsertAtEnding(data string) {
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

// Length returns the number of nodes in the linked list.
func (l *Adapter) Length() int {
	var len = 0
	ite := l.head
	for ite != nil {
		ite = ite.Next
		len = len + 1
	}
	return len
}

// DelAtBeg deletes the first node in the linked list and returns a success message.
func (l *Adapter) DelAtBeg() string {
	if l.head == nil {
		return "no elements in the linked list"
	} else {
		l.head = l.head.Next
		return "first element is deleted successfully"
	}
}

// DelAtEnd deletes the last node in the linked list and returns a success message.
func (l *Adapter) DelAtEnd() string {
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

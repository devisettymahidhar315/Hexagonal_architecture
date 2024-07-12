package usecase

import "app/adapter"

// Usecase represents the use case layer that interacts with the adapter.
type Usecase struct {
	repo *adapter.Adapter
}

// Init initializes a new Usecase instance with a given adapter.
func Init(j *adapter.Adapter) *Usecase {
	return &Usecase{repo: j}
}

// Print retrieves and returns all elements in the linked list as a concatenated string.
func (s *Usecase) Print() string {
	print := s.repo.Print()
	return print
}

// Insertatstarting inserts a new element at the beginning of the linked list.
func (s *Usecase) Insertatstarting(s1 string) {
	s.repo.InsertAtBeginning(s1)
}

// InsertAtEnding inserts a new element at the end of the linked list.
func (s *Usecase) InsertAtEnding(s1 string) {
	s.repo.InsertAtEnding(s1)
}

// Length retrieves and returns the number of elements in the linked list.
func (s *Usecase) Length() int {
	return s.repo.Length()
}

// DelAtBeg deletes the first element in the linked list and returns a success message.
func (s *Usecase) DelAtBeg() string {
	return s.repo.DelAtBeg()
}

// DelAtEnd deletes the last element in the linked list and returns a success message.
func (s *Usecase) DelAtEnd() string {
	return s.repo.DelAtEnd()
}

package usecase

import "app/adapter"

type K struct {
	sd *adapter.JK
}

func Init(j *adapter.JK) *K {
	return &K{sd: j}
}

func (s *K) Print() string {
	e := s.sd.Print()
	return e
}

func (s *K) Insertatstarting(s1 string) {
	s.sd.InsertAtBeginning(s1)

}

func (s *K) InsertAtEnding(s1 string) {
	s.sd.InsertAtEnding(s1)
}

func (s *K) Length() int {
	return s.sd.Length()
}

func (s *K) DelAtBeg() string {
	return s.sd.DelAtBeg()
}

func (s *K) DelAtEnd() string {
	return s.sd.DelAtEnd()
}

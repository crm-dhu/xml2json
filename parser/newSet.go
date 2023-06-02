package parser

var Exists = struct{}{}

type Set struct {
	m map[any]struct{}
}

func NewSet(items ...any) *Set {
	s := &Set{}
	s.m = make(map[any]struct{})
	s.Add(items...)
	return s
}

func (s *Set) Add(items ...any) error {
	for _, item := range items {
		s.m[item] = Exists
	}
	return nil
}

func (s *Set) AddSet(other *Set) error {
	for key := range other.m {
		s.Add(key)
	}
	return nil
}

func (s *Set) Contains(item any) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Clear() {
	s.m = make(map[any]struct{})
}

func (s *Set) Equal(other *Set) bool {
	if s.Size() != other.Size() {
		return false
	}

	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func (s *Set) IsSubset(other *Set) bool {
	if s.Size() > other.Size() {
		return false
	}
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

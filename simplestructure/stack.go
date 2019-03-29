package simplestructure

import (
	"container/list"
	"sync"
)

type Stack interface {
	Pushback(interface{})
	Popback() interface{}
	Size() int
}

type SimpleStack struct {
	elements *list.List
	lock     sync.Mutex
}

func NewSimpleStack() *SimpleStack {
	l := list.New()

	s := SimpleStack{
		elements: l,
	}

	return &s
}

func (s *SimpleStack) Pushback(v interface{}) {
	defer s.lock.Unlock()

	s.lock.Lock()

	s.elements.PushBack(v)
}

func (s *SimpleStack) Popback() interface{} {
	defer s.lock.Unlock()

	s.lock.Lock()

	e := s.elements.Back()
	if e == nil {
		return nil
	}
	return s.elements.Remove(e)
}

func (s *SimpleStack) Size() int {
	return s.elements.Len()
}

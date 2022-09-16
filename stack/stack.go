package stack

import (
	"fmt"
	"sync"
)

// Stack represent a stack
type Stack struct {
	lock sync.RWMutex
	data []interface{}
}

//	ErrorEmpty occurred on illegal operations on empty queue
var ErrorEmpty = fmt.Errorf("empty queue")

//	Len returned length of stack
func (s *Stack) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.data)
}

// Push adds an item to the stack
func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data = append(s.data, item)
}

//	Peek returns the top item from the stack without removing it
func (s *Stack) Peek() (interface{}, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if len(s.data) == 0 {
		return nil, ErrorEmpty
	}

	return s.data[len(s.data)-1], nil
}

//	Pop returns the top item from the stack
func (s *Stack) Pop() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.data) == 0 {
		return nil, ErrorEmpty
	}

	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return item, nil
}

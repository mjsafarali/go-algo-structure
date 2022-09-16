package queue

import (
	"fmt"
	"sync"
)

// Queue represent a queue
type Queue struct {
	lock sync.RWMutex
	data []interface{}
}

//	ErrorEmpty occurred on illegal operations on empty queue
var ErrorEmpty = fmt.Errorf("empty queue")

//	Len returned length of queue
func (q *Queue) Len() int {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return len(q.data)
}

//	Add will adds an item to queue
func (q *Queue) Add(item interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.data = append(q.data, item)
}

// Peek will return first item in the queue
func (q *Queue) Peek() (interface{}, error) {
	q.lock.RLock()
	defer q.lock.RUnlock()

	if len(q.data) == 0 {
		return nil, ErrorEmpty
	}

	return q.data[0], nil
}

// Remove will detach first item of queue
func (q *Queue) Remove() error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.data) == 0 {
		return ErrorEmpty
	}

	q.data = q.data[1:]

	return nil
}

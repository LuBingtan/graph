package simplestructure

import (
	"container/list"
	"sync"
)

/**********************************************************************************/
// define queue interface
/**********************************************************************************/

type Queue interface {
	Pushback(interface{})
	Popfront() interface{}
	Size() int
}

/**********************************************************************************/
// define simple queue
/**********************************************************************************/

type SimpleQueue struct {
	elements *list.List
	lock     sync.RWMutex
}

func NewSimpleQueue() *SimpleQueue {
	l := list.New()

	q := SimpleQueue{
		elements: l,
	}

	return &q
}

func (q *SimpleQueue) Pushback(v interface{}) {
	defer q.lock.Unlock()
	q.lock.Lock()

	q.elements.PushBack(v)
}

func (q *SimpleQueue) Popfront() interface{} {
	defer q.lock.Unlock()
	q.lock.Lock()

	e := q.elements.Front()
	if e == nil {
		return nil
	}
	return q.elements.Remove(e)
}

func (q *SimpleQueue) Size() int {
	defer q.lock.RUnlock()
	q.lock.RLock()

	return q.elements.Len()
}

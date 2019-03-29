package simplestructure

import (
	"sync"
	"reflect"
)

/***************************************** define vector interface *****************************************/

type Vector interface {
	Pushback(interface{})
	Insert(interface{}, int)

	Remove(int) interface{}
	Popback() interface{}
	Popfront() interface{}

	At(int) interface{}
	Len() int
	Iterator() Iterator
}

/***************************************** define vector iterator *****************************************/

type vectorIterator struct {
	vec   Vector
	index int
}

func (it *vectorIterator) HasNext() bool {
	return it.index <= it.vec.Len()-1
}

func (it *vectorIterator) Next() interface{} {
	if !it.HasNext() {
		return nil
	}

	it.index++
	return it.vec.At(it.index - 1)
}

/***************************************** define simple vector *****************************************/

type SimpleVector struct {
	data []interface{}
	lock sync.RWMutex
}

func NewSimpleVector() *SimpleVector {
	vec := SimpleVector{}
	return &vec
}

func (vec *SimpleVector) Pushback(v interface{}) {
	defer vec.lock.Unlock()

	vec.lock.Lock()

	vec.data = append(vec.data, v)
}

func (vec *SimpleVector) Insert(v interface{}, next int) {
	defer vec.lock.Unlock()

	vec.lock.Lock()

	vec.data = append(vec.data, nil)
	copy(vec.data[next:], vec.data[next+1:len(vec.data)])
	vec.data[next] = v
}

func (vec *SimpleVector) Remove(index int) interface{} {
	defer vec.lock.Unlock()

	vec.lock.Lock()

	if index > len(vec.data)-1 {
		return nil
	}

	v := vec.data[index]
	vec.data = append(vec.data[:index], vec.data[index+1:]...)

	return v
}

func (vec *SimpleVector) Popback() interface{} {
	return vec.Remove(len(vec.data) - 1)
}

func (vec *SimpleVector) Popfront() interface{} {
	return vec.Remove(0)
}

func (vec *SimpleVector) At(index int) interface{} {
	defer vec.lock.RUnlock()

	vec.lock.RLock()

	if index > len(vec.data)-1 {
		return nil
	}

	return vec.data[index]
}

func (vec *SimpleVector) Find(v interface{}) int {
	for i, d := range vec.data {
		if reflect.DeepEqual(v, d) {
			return i
		}
	}

	return -1
}

func (vec *SimpleVector) Len() int {
	defer vec.lock.RUnlock()

	vec.lock.RLock()

	return len(vec.data)
}

func (vec *SimpleVector) Iterator() Iterator {
	return &vectorIterator{
		vec:   vec,
		index: 0,
	}
}

func (vec *SimpleVector) Data() []interface{} {
	return vec.data
}
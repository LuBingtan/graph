package simplestructure

import (
	"testing"
)

func PrintIterator(t *testing.T, it Iterator) {
	t.Log("==========PrintIterator start===========")
	for n := it.Next(); n != nil; n = it.Next() {
		t.Log(n)
	}
	t.Log("==========PrintIterator end===========")
}

func Test4Vector4Create(t *testing.T) {
	vec := NewSimpleVector()
	vec.Pushback(0)
	vec.Pushback(1)
	vec.Pushback(2)
	vec.Pushback(3)
	vec.Pushback(4)
	t.Log(vec.Len())
	PrintIterator(t, vec.Iterator())

	t.Logf("pop back:%v", vec.Popback())
	PrintIterator(t, vec.Iterator())

	t.Logf("pop front:%v", vec.Popfront())
	PrintIterator(t, vec.Iterator())

	t.Logf("remove:%v", vec.Remove(3))
	PrintIterator(t, vec.Iterator())
}

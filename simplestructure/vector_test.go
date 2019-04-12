package simplestructure

import (
	"testing"
)

func PrintData(t *testing.T, l []interface{}) {
	t.Log("==========PrintData start===========")
	for _, n := range l {
		t.Log(n)
	}
	t.Log("==========PrintData end===========")
}

func Test4Vector4Create(t *testing.T) {
	vec := NewSimpleVector()
	vec.Pushback(0)
	vec.Pushback(1)
	vec.Pushback(2)
	vec.Pushback(3)
	vec.Pushback(4)
	t.Log(vec.Len())
	PrintData(t, vec.Data())

	t.Logf("pop back:%v", vec.Popback())
	PrintData(t, vec.Data())

	t.Logf("pop front:%v", vec.Popfront())
	PrintData(t, vec.Data())

	t.Logf("remove:%v", vec.Remove(3))
	PrintData(t, vec.Data())
}

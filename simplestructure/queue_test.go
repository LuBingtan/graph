package simplestructure

import "testing"

func Test4Queue(t *testing.T) {
	q := NewSimpleQueue()
	q.Pushback(0)
	q.Pushback(1)
	q.Pushback(2)

	for {
		v := q.Popfront()
		if v == nil {
			break
		}
		t.Log(v)
	}
}

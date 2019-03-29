package simplestructure

import "testing"

func Test4Stack(t *testing.T) {
	q := NewSimpleStack()
	q.Pushback(0)
	q.Pushback(1)
	q.Pushback(2)

	for {
		v := q.Popback()
		if v == nil {
			break
		}
		t.Log(v)
	}
}

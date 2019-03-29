package simplestructure

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

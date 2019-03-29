package graph

import "sync"

/*****************************************  edge interface  *****************************************/

// define for edge type
type EdgeType string

// define for edge state
type EdgeState string

// define for edge interface
type EdgeInterface interface {
	/////// meta data ///////
	// update
	SetType(EdgeType)
	SetWeight(int)
	SetState(EdgeState)
	// read
	Type() EdgeType
	Weight() int
	State() EdgeState

	/////// relation data ///////
	// update
	SetVertex(VertexInterface)
	// read
	Vertex() VertexInterface
}

/*****************************************  edge struct  *****************************************/

type AbstractEdge struct {
	// meta data
	edgeType EdgeType
	weight   int
	state    EdgeState
	// graph data
	vertex VertexInterface
	// mutex
	mutex sync.RWMutex
}

func NewEdge() *AbstractEdge {
	return &AbstractEdge{}
}

func (e *AbstractEdge) SetType(t EdgeType) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	e.edgeType = t
}

func (e *AbstractEdge) SetWeight(w int) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	e.weight = w
}

func (e *AbstractEdge) SetState(s EdgeState) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	e.state = s
}

func (e *AbstractEdge) Type() EdgeType {
	defer e.mutex.RUnlock()
	e.mutex.RLock()
	return e.edgeType
}

func (e *AbstractEdge) Weight() int {
	defer e.mutex.RUnlock()
	e.mutex.RLock()
	return e.weight
}

func (e *AbstractEdge) State() EdgeState {
	defer e.mutex.RUnlock()
	e.mutex.RLock()
	return e.state
}

func (e *AbstractEdge) SetVertex(v VertexInterface) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	e.vertex = v
}

func (e *AbstractEdge) Vertex() VertexInterface {
	defer e.mutex.RUnlock()
	e.mutex.RLock()
	return e.vertex
}

package graph

import "testing"

func TestInitTraversalState(t *testing.T) {

	var count = 500

	g := CreateGraph(false)

	for i := 0; i < count; i++ {
		g.InsertEdge(i, i+1, false)
	}

	state := InitTraversalState(g)

	if len(state.discovered) < 500 {
		t.Errorf("Incorrect state field size:%d", len(state.discovered))
	}

}

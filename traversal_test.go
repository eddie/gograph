package graph

import "testing"

var graph_undirected = [][]int{{1, 6}, {1, 2}, {1, 5}, {2, 5}, {2, 3}, {3, 4}, {4, 5}}
var undirected_parents = []int{0, -1, 1, 2, 5, 1, 1}

func GetTestGraphUndirected() *Graph {

	g := CreateGraph(false)

	for _, i := range graph_undirected {

		g.InsertEdge(i[0], i[1], false)
	}

	return g
}

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

func TestBFSParentRealtion(t *testing.T) {

	g := GetTestGraphUndirected()

	funcs := &TraversalFuncs{}
	state := InitTraversalState(g)

	g.BFS(1, funcs, state)

	for i, p := range undirected_parents {

		if p != state.parents[i] {
			t.Errorf("Parent relation failed %d,%d", p, state.parents[i])
		}
	}

}

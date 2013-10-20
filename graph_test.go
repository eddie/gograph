package graph

import "testing"

func TestCreateGraph(t *testing.T) {

	g := CreateGraph(false)

	g.edges[0] = &Edgenode{}

	if g.edges[0] == nil {
		t.Errorf("Graph edge not set")
	}
	if g.directed != false {
		t.Errorf("Graph directed state")
	}
}

func TestInsertEdge(t *testing.T) {

	var count = 500

	g := CreateGraph(false)

	for i := 0; i < count; i++ {
		g.InsertEdge(i, i+1, false)
	}

	for i := 0; i < count; i++ {
		if g.edges[i].y != i+1 {
			t.Errorf("Edge not inserted")
		}

		if g.degree[i] > 2 {
			t.Errorf("Degree incorrect: %d", g.degree[i])
		}
	}

}

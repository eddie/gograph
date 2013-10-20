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

func TestUndirectedEdge(t *testing.T) {

	g := CreateGraph(false)

	g.InsertEdge(0, 1, false)

	if g.edges[0].y != 1 {
		t.Errorf("Edge not created")
	}
	if g.edges[1].y != 0 {
		t.Errorf("Edge not created")
	}
}

func TestDirectedEdge(t *testing.T) {

	g := CreateGraph(true)

	g.InsertEdge(0, 1, true)

	if g.edges[0].y != 1 {
		t.Errorf("Edge not created")
	}

	if g.edges[1] != nil {
		t.Errorf("Edge should not be created (Directed Edge)")
	}
}

func BenchmarkInsertEdge(b *testing.B) {

	var count = 1000000

	g := CreateGraph(false)

	for i := 0; i < count; i++ {
		g.InsertEdge(i, i+1, false)
	}
}

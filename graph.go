// Basic Graph Adjacency List implementation and various operations
// Correctness or efficiency not guaranteed. Use at your own risk.
//
// Eddie Blundell - eblundell@gmail.com
// (Mostly transLated from The Algorithm Design Manual - Steven S. Skiena) 

package graph

import (
	"bufio"
	"fmt"
	"os"
)

const MAXVERT = 100

type Edgenode struct {
	y      int
	weight int
	next   *Edgenode
}

// TODO: merge processed,discovered to enum type
type Graph struct {
	edges     []*Edgenode
	degree    []int
	nvertices int
	nedges    int
	directed  bool
}

func CreateGraph(directed bool) (g *Graph) {

	g = &Graph{
		make([]*Edgenode, 10, MAXVERT),
		make([]int, 10, MAXVERT),
		0,
		0,
		directed,
	}

	return g
}

func (g *Graph) InsertEdge(x, y int, directed bool) {

	edgenode := &Edgenode{y, 0, nil}
	edgenode.next = g.edges[x]

	g.edges[x] = edgenode
	g.degree[x]++

	if !directed {
		g.InsertEdge(y, x, true)
	} else {
		g.nedges++
	}
}

func (g *Graph) ReadGraph(fname string) {

	file, err := os.Open(fname)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	var x, y int

	for {
		n, err := fmt.Fscanf(reader, "%d %d\n", &x, &y)

		if err != nil {
			break
		}

		if n <= 0 {
			continue
		}

		g.InsertEdge(x, y, g.directed)
		g.nvertices++
	}
}

func (g *Graph) PrintGraph() {
	for i := 1; i <= g.nvertices; i++ {

		var p *Edgenode
		p = g.edges[i]

		if p != nil {
			fmt.Printf("Vertex: %d Adjency Vertices: ", i)
		}
		for p != nil {

			fmt.Printf("%d ", p.y)
			p = p.next
		}

		fmt.Printf("\n")
	}
}

func (g *Graph) VertexCount() int {
	return g.nvertices
}

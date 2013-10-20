package graph

type TraversalState struct {
	processed  []bool
	discovered []bool
	entry_time []int
	exit_time  []int
	parents    []int
	time       int
	finished   bool
}

type TraversalFuncs struct {
	Early func(int, *TraversalState)
	Late  func(int, *TraversalState)
	Edge  func(int, int, *TraversalState)
}

func (f *TraversalFuncs) ProcessEarly(v int, s *TraversalState) {
	if f.Early != nil {
		f.Early(v, s)
	}
}

func (f *TraversalFuncs) ProcessLate(v int, s *TraversalState) {
	if f.Late != nil {
		f.Late(v, s)
	}
}

func (f *TraversalFuncs) ProcessEdge(x, y int, s *TraversalState) {
	if f.Edge != nil {
		f.Edge(x, y, s)
	}
}

func InitTraversalState() (s *TraversalState) {

	return &TraversalState{
		make([]bool, 10, MAXVERT),
		make([]bool, 10, MAXVERT),
		make([]int, 10, MAXVERT),
		make([]int, 10, MAXVERT),
		make([]int, 10, MAXVERT),
		0,
		false,
	}
}

func (g *Graph) BFS(start int, funcs *TraversalFuncs, state *TraversalState) {

	if state == nil {
		state = InitTraversalState()
	}

	q := NewQueue(20)
	q.Push(Node(start))

	state.processed[start] = true
	state.parents[start] = -1

	for q.Empty() == false {

		v := int(q.Pop())

		state.discovered[v] = true
		p := g.edges[v]

		funcs.ProcessEarly(v, state)

		for p != nil {

			y := p.y

			if state.processed[y] == false || g.directed {
				funcs.ProcessEdge(v, y, state)
			}

			if state.discovered[y] == false {
				q.Push(Node(y))
				state.discovered[y] = true
				state.parents[y] = v
			}

			p = p.next
		}

		funcs.ProcessLate(v, state)
	}
}

func (g *Graph) DFS(v int, funcs *TraversalFuncs, state *TraversalState) {

	if state == nil {
		state = InitTraversalState()
	}

	if state.finished {
		return
	}

	state.discovered[v] = true
	state.time++
	state.entry_time[v] = state.time

	funcs.ProcessEarly(v, state)

	p := g.edges[v]
	for p != nil {

		y := p.y
		if state.discovered[y] == false {

			state.parents[y] = v
			funcs.ProcessEdge(v, y, state)
			g.DFS(y, funcs, state)

		} else if (!state.processed[y] && state.parents[v] != y) || g.directed {
			funcs.ProcessEdge(v, y, state)
		}

		if state.finished {
			return
		}

		p = p.next
	}

	funcs.ProcessLate(v, state)

	state.time++
	state.exit_time[v] = state.time
	state.processed[v] = true
}

func (g *Graph) FindPathExt(start, end int, cb func(int), parents []int) {

	if (start == end) || (end == -1) {
		cb(start)
	} else {
		g.FindPathExt(start, parents[end], cb, parents)
		cb(end)
	}
}

func (g *Graph) FindPath(start, end int, cb func(int)) {

	var state *TraversalState = InitTraversalState()
	var funcs *TraversalFuncs = &TraversalFuncs{}

	g.BFS(start, funcs, state)
	g.FindPathExt(start, end, cb, state.parents)
}

func (s *TraversalState) Discovered(v int) bool {
	return s.discovered[v]
}

func (s *TraversalState) Parent(v int) int {
	return s.parents[v]
}

func (s *TraversalState) Finished(finished bool) {
	s.finished = finished
}

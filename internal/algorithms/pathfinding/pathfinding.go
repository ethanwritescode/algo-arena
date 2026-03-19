package pathfinding

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand/v2"
)

// Cell represents a cell in the grid
type Cell struct {
	Row, Col int
}

// Grid represents the pathfinding grid
type Grid struct {
	Width  int
	Height int
	Walls  map[Cell]bool
	Start  Cell
	End    Cell
}

// Step represents a single step in the pathfinding visualization
type Step struct {
	Grid        *Grid
	Visited     map[Cell]bool
	Current     Cell
	Frontier    []Cell
	Path        []Cell
	Description string
}

// Algorithm represents a pathfinding algorithm
type Algorithm struct {
	Name        string
	Description string
	TimeComplex string
	Steps       []Step
}

// NewGrid creates a new grid with a proper maze
func NewGrid(width, height int, wallDensity float64) *Grid {
	g := &Grid{
		Width:  width,
		Height: height,
		Walls:  make(map[Cell]bool),
		Start:  Cell{Row: 1, Col: 1},
		End:    Cell{Row: height - 2, Col: width - 2},
	}

	// Generate a proper maze using recursive backtracking
	g.generateRecursiveBacktrackingMaze()

	return g
}

// generateRecursiveBacktrackingMaze creates a proper maze with corridors
func (g *Grid) generateRecursiveBacktrackingMaze() {
	// Start with all walls
	for row := 0; row < g.Height; row++ {
		for col := 0; col < g.Width; col++ {
			g.Walls[Cell{Row: row, Col: col}] = true
		}
	}

	// Carve out the maze using recursive backtracking
	// Start from an odd cell to ensure proper maze structure
	startRow := 1
	startCol := 1

	// Mark start as passage
	delete(g.Walls, Cell{Row: startRow, Col: startCol})

	// Stack for backtracking
	stack := []Cell{{Row: startRow, Col: startCol}}
	visited := make(map[Cell]bool)
	visited[Cell{Row: startRow, Col: startCol}] = true

	// Directions: up, down, left, right (move by 2 to leave walls between passages)
	directions := []Cell{
		{Row: -2, Col: 0},
		{Row: 2, Col: 0},
		{Row: 0, Col: -2},
		{Row: 0, Col: 2},
	}

	for len(stack) > 0 {
		current := stack[len(stack)-1]

		// Find unvisited neighbors
		var unvisitedNeighbors []Cell
		for _, dir := range directions {
			next := Cell{Row: current.Row + dir.Row, Col: current.Col + dir.Col}
			// Check bounds (must be odd coordinates and within maze)
			if next.Row > 0 && next.Row < g.Height-1 &&
				next.Col > 0 && next.Col < g.Width-1 &&
				!visited[next] {
				unvisitedNeighbors = append(unvisitedNeighbors, next)
			}
		}

		if len(unvisitedNeighbors) > 0 {
			// Choose random neighbor
			next := unvisitedNeighbors[rand.IntN(len(unvisitedNeighbors))]
			visited[next] = true

			// Remove wall between current and next
			wallRow := current.Row + (next.Row-current.Row)/2
			wallCol := current.Col + (next.Col-current.Col)/2
			delete(g.Walls, Cell{Row: wallRow, Col: wallCol})
			delete(g.Walls, next)

			stack = append(stack, next)
		} else {
			// Backtrack
			stack = stack[:len(stack)-1]
		}
	}

	// Ensure start and end are accessible
	delete(g.Walls, g.Start)
	delete(g.Walls, g.End)

	// Clear cells around start and end for better accessibility
	for _, dir := range []Cell{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		startNeighbor := Cell{Row: g.Start.Row + dir.Row, Col: g.Start.Col + dir.Col}
		endNeighbor := Cell{Row: g.End.Row + dir.Row, Col: g.End.Col + dir.Col}
		if startNeighbor.Row > 0 && startNeighbor.Row < g.Height-1 &&
			startNeighbor.Col > 0 && startNeighbor.Col < g.Width-1 {
			delete(g.Walls, startNeighbor)
		}
		if endNeighbor.Row > 0 && endNeighbor.Row < g.Height-1 &&
			endNeighbor.Col > 0 && endNeighbor.Col < g.Width-1 {
			delete(g.Walls, endNeighbor)
		}
	}

	// Add some random passages to create multiple path options (makes it more interesting)
	extraPassages := (g.Width * g.Height) / 20
	for i := 0; i < extraPassages; i++ {
		row := rand.IntN(g.Height-2) + 1
		col := rand.IntN(g.Width-2) + 1
		cell := Cell{Row: row, Col: col}
		if cell != g.Start && cell != g.End {
			delete(g.Walls, cell)
		}
	}
}

// GetNeighbors returns valid neighboring cells
func (g *Grid) GetNeighbors(c Cell) []Cell {
	directions := []Cell{
		{Row: -1, Col: 0}, // Up
		{Row: 1, Col: 0},  // Down
		{Row: 0, Col: -1}, // Left
		{Row: 0, Col: 1},  // Right
	}

	neighbors := []Cell{}
	for _, d := range directions {
		n := Cell{Row: c.Row + d.Row, Col: c.Col + d.Col}
		if n.Row >= 0 && n.Row < g.Height && n.Col >= 0 && n.Col < g.Width {
			if !g.Walls[n] {
				neighbors = append(neighbors, n)
			}
		}
	}
	return neighbors
}

// Manhattan distance heuristic
func heuristic(a, b Cell) float64 {
	return math.Abs(float64(a.Row-b.Row)) + math.Abs(float64(a.Col-b.Col))
}

// copyVisited creates a copy of the visited map
func copyVisited(v map[Cell]bool) map[Cell]bool {
	result := make(map[Cell]bool)
	for k, val := range v {
		result[k] = val
	}
	return result
}

// copyCells creates a copy of a cell slice
func copyCells(cells []Cell) []Cell {
	result := make([]Cell, len(cells))
	copy(result, cells)
	return result
}

// BFS performs breadth-first search
func BFS(grid *Grid) *Algorithm {
	algo := &Algorithm{
		Name:        "Breadth-First Search (BFS)",
		Description: "Explores level by level in expanding wavefront — guarantees shortest path in unweighted graphs",
		TimeComplex: "O(V + E)",
		Steps:       []Step{},
	}

	visited := make(map[Cell]bool)
	parent := make(map[Cell]Cell)
	queue := []Cell{grid.Start}
	visited[grid.Start] = true

	algo.Steps = append(algo.Steps, Step{
		Grid:        grid,
		Visited:     copyVisited(visited),
		Current:     grid.Start,
		Frontier:    copyCells(queue),
		Description: fmt.Sprintf("Start BFS at (%d,%d) → target (%d,%d)", grid.Start.Row, grid.Start.Col, grid.End.Row, grid.End.Col),
	})

	found := false
	for len(queue) > 0 && !found {
		// Process entire current level for cleaner visualization
		levelSize := len(queue)
		for i := 0; i < levelSize && !found; i++ {
			current := queue[0]
			queue = queue[1:]

			if current == grid.End {
				found = true
				algo.Steps = append(algo.Steps, Step{
					Grid:        grid,
					Visited:     copyVisited(visited),
					Current:     current,
					Frontier:    copyCells(queue),
					Description: "Found target!",
				})
				break
			}

			for _, neighbor := range grid.GetNeighbors(current) {
				if !visited[neighbor] {
					visited[neighbor] = true
					parent[neighbor] = current
					queue = append(queue, neighbor)
				}
			}
		}

		if !found && len(queue) > 0 {
			algo.Steps = append(algo.Steps, Step{
				Grid:        grid,
				Visited:     copyVisited(visited),
				Current:     queue[0],
				Frontier:    copyCells(queue),
				Description: fmt.Sprintf("Expanding wave — visited: %d, frontier: %d", len(visited), len(queue)),
			})
		}
	}

	// Reconstruct path
	path := reconstructPath(parent, grid.Start, grid.End, found)

	algo.Steps = append(algo.Steps, Step{
		Grid:        grid,
		Visited:     copyVisited(visited),
		Path:        path,
		Description: getResultDescription(found, len(path), len(visited)),
	})

	return algo
}

// DFS performs depth-first search
func DFS(grid *Grid) *Algorithm {
	algo := &Algorithm{
		Name:        "Depth-First Search (DFS)",
		Description: "Explores deeply along each branch before backtracking — memory efficient but may find longer paths",
		TimeComplex: "O(V + E)",
		Steps:       []Step{},
	}

	visited := make(map[Cell]bool)
	parent := make(map[Cell]Cell)
	stack := []Cell{grid.Start}

	algo.Steps = append(algo.Steps, Step{
		Grid:        grid,
		Visited:     copyVisited(visited),
		Current:     grid.Start,
		Frontier:    copyCells(stack),
		Description: fmt.Sprintf("Start DFS at (%d,%d) → target (%d,%d)", grid.Start.Row, grid.Start.Col, grid.End.Row, grid.End.Col),
	})

	found := false
	stepCounter := 0
	for len(stack) > 0 && !found {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[current] {
			continue
		}
		visited[current] = true
		stepCounter++

		if current == grid.End {
			found = true
			algo.Steps = append(algo.Steps, Step{
				Grid:        grid,
				Visited:     copyVisited(visited),
				Current:     current,
				Frontier:    copyCells(stack),
				Description: "Found target!",
			})
			break
		}

		// Add unvisited neighbors to stack (reverse order for consistent direction preference)
		neighbors := grid.GetNeighbors(current)
		addedAny := false
		for i := len(neighbors) - 1; i >= 0; i-- {
			neighbor := neighbors[i]
			if !visited[neighbor] {
				parent[neighbor] = current
				stack = append(stack, neighbor)
				addedAny = true
			}
		}

		// Record step periodically or at interesting points
		if stepCounter%3 == 0 || !addedAny {
			desc := fmt.Sprintf("Exploring (%d,%d) — depth: %d", current.Row, current.Col, len(stack))
			if !addedAny && len(stack) > 0 {
				desc = fmt.Sprintf("Dead end at (%d,%d) — backtracking", current.Row, current.Col)
			}
			algo.Steps = append(algo.Steps, Step{
				Grid:        grid,
				Visited:     copyVisited(visited),
				Current:     current,
				Frontier:    copyCells(stack),
				Description: desc,
			})
		}
	}

	// Reconstruct path
	path := reconstructPath(parent, grid.Start, grid.End, found)

	algo.Steps = append(algo.Steps, Step{
		Grid:        grid,
		Visited:     copyVisited(visited),
		Path:        path,
		Description: getResultDescription(found, len(path), len(visited)),
	})

	return algo
}

// Priority queue implementation for A* and Dijkstra
type pqItem struct {
	cell     Cell
	priority float64
	index    int
}

type priorityQueue []*pqItem

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pqItem)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// Dijkstra performs Dijkstra's algorithm
func Dijkstra(grid *Grid) *Algorithm {
	algo := &Algorithm{
		Name:        "Dijkstra's Algorithm",
		Description: "Expands outward by shortest distance first — optimal for weighted graphs, equivalent to BFS for unit weights",
		TimeComplex: "O((V + E) log V)",
		Steps:       []Step{},
	}

	visited := make(map[Cell]bool)
	dist := make(map[Cell]float64)
	parent := make(map[Cell]Cell)

	// Initialize distances
	for row := 0; row < grid.Height; row++ {
		for col := 0; col < grid.Width; col++ {
			dist[Cell{Row: row, Col: col}] = math.Inf(1)
		}
	}
	dist[grid.Start] = 0

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &pqItem{cell: grid.Start, priority: 0})

	algo.Steps = append(algo.Steps, Step{
		Grid:        grid,
		Visited:     copyVisited(visited),
		Current:     grid.Start,
		Frontier:    []Cell{grid.Start},
		Description: fmt.Sprintf("Start Dijkstra at (%d,%d) — all distances = ∞ except start = 0", grid.Start.Row, grid.Start.Col),
	})

	found := false
	stepCounter := 0
	for pq.Len() > 0 && !found {
		item := heap.Pop(pq).(*pqItem)
		current := item.cell
		currentDist := item.priority

		if visited[current] {
			continue
		}
		visited[current] = true
		stepCounter++

		if current == grid.End {
			found = true
			algo.Steps = append(algo.Steps, Step{
				Grid:        grid,
				Visited:     copyVisited(visited),
				Current:     current,
				Frontier:    getFrontierCells(pq),
				Description: fmt.Sprintf("Found target! Shortest distance: %.0f", currentDist),
			})
			break
		}

		// Update neighbors
		for _, neighbor := range grid.GetNeighbors(current) {
			if !visited[neighbor] {
				newDist := dist[current] + 1
				if newDist < dist[neighbor] {
					dist[neighbor] = newDist
					parent[neighbor] = current
					heap.Push(pq, &pqItem{cell: neighbor, priority: newDist})
				}
			}
		}

		// Record step periodically
		if stepCounter%4 == 0 {
			algo.Steps = append(algo.Steps, Step{
				Grid:        grid,
				Visited:     copyVisited(visited),
				Current:     current,
				Frontier:    getFrontierCells(pq),
				Description: fmt.Sprintf("Processing dist=%.0f — visited: %d, queue: %d", currentDist, len(visited), pq.Len()),
			})
		}
	}

	// Reconstruct path
	path := reconstructPath(parent, grid.Start, grid.End, found)

	algo.Steps = append(algo.Steps, Step{
		Grid:        grid,
		Visited:     copyVisited(visited),
		Path:        path,
		Description: getResultDescription(found, len(path), len(visited)),
	})

	return algo
}

// getFrontierCells extracts cells from the priority queue for visualization
func getFrontierCells(pq *priorityQueue) []Cell {
	seen := make(map[Cell]bool)
	cells := make([]Cell, 0, pq.Len())
	for _, item := range *pq {
		if !seen[item.cell] {
			seen[item.cell] = true
			cells = append(cells, item.cell)
		}
	}
	return cells
}

// AStar performs A* pathfinding algorithm
func AStar(grid *Grid) *Algorithm {
	algo := &Algorithm{
		Name:        "A* Search Algorithm",
		Description: "Uses heuristic to guide search toward goal — optimal and typically much faster than Dijkstra",
		TimeComplex: "O(E log V)",
		Steps:       []Step{},
	}

	visited := make(map[Cell]bool)
	gScore := make(map[Cell]float64) // Cost from start to node
	fScore := make(map[Cell]float64) // gScore + heuristic estimate to goal
	parent := make(map[Cell]Cell)

	// Initialize scores
	for row := 0; row < grid.Height; row++ {
		for col := 0; col < grid.Width; col++ {
			gScore[Cell{Row: row, Col: col}] = math.Inf(1)
			fScore[Cell{Row: row, Col: col}] = math.Inf(1)
		}
	}
	gScore[grid.Start] = 0
	fScore[grid.Start] = heuristic(grid.Start, grid.End)

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &pqItem{cell: grid.Start, priority: fScore[grid.Start]})

	algo.Steps = append(algo.Steps, Step{
		Grid:        grid,
		Visited:     copyVisited(visited),
		Current:     grid.Start,
		Frontier:    []Cell{grid.Start},
		Description: fmt.Sprintf("Start A* — heuristic to goal: %.0f (Manhattan distance)", fScore[grid.Start]),
	})

	found := false
	stepCounter := 0
	for pq.Len() > 0 && !found {
		item := heap.Pop(pq).(*pqItem)
		current := item.cell

		if visited[current] {
			continue
		}
		visited[current] = true
		stepCounter++

		g := gScore[current]
		h := heuristic(current, grid.End)

		if current == grid.End {
			found = true
			algo.Steps = append(algo.Steps, Step{
				Grid:        grid,
				Visited:     copyVisited(visited),
				Current:     current,
				Frontier:    getFrontierCells(pq),
				Description: fmt.Sprintf("Found target! Path cost: %.0f", g),
			})
			break
		}

		// Expand neighbors
		for _, neighbor := range grid.GetNeighbors(current) {
			if visited[neighbor] {
				continue
			}

			tentativeG := gScore[current] + 1
			if tentativeG < gScore[neighbor] {
				parent[neighbor] = current
				gScore[neighbor] = tentativeG
				fScore[neighbor] = tentativeG + heuristic(neighbor, grid.End)
				heap.Push(pq, &pqItem{cell: neighbor, priority: fScore[neighbor]})
			}
		}

		// Record step periodically
		if stepCounter%3 == 0 {
			algo.Steps = append(algo.Steps, Step{
				Grid:        grid,
				Visited:     copyVisited(visited),
				Current:     current,
				Frontier:    getFrontierCells(pq),
				Description: fmt.Sprintf("f=%.0f (g=%.0f + h=%.0f) — visited: %d, open: %d", g+h, g, h, len(visited), pq.Len()),
			})
		}
	}

	// Reconstruct path
	path := reconstructPath(parent, grid.Start, grid.End, found)

	algo.Steps = append(algo.Steps, Step{
		Grid:        grid,
		Visited:     copyVisited(visited),
		Path:        path,
		Description: getResultDescription(found, len(path), len(visited)),
	})

	return algo
}

// reconstructPath builds the path from parent map
func reconstructPath(parent map[Cell]Cell, start, end Cell, found bool) []Cell {
	if !found {
		return []Cell{}
	}
	path := []Cell{}
	current := end
	for current != start {
		path = append([]Cell{current}, path...)
		var ok bool
		current, ok = parent[current]
		if !ok {
			return []Cell{} // Path reconstruction failed
		}
	}
	path = append([]Cell{start}, path...)
	return path
}

func getResultDescription(found bool, pathLen int, visitedCount int) string {
	if found {
		return fmt.Sprintf("Path found! Length: %d, visited: %d cells (%.0f%% of search space explored)",
			pathLen, visitedCount, float64(visitedCount)/float64(pathLen)*100/5)
	}
	return fmt.Sprintf("No path exists! Explored %d cells", visitedCount)
}

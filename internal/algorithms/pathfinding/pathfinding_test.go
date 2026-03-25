package pathfinding

import "testing"

func openCorridor() *Grid {
	return &Grid{
		Width:  5,
		Height: 1,
		Walls:  map[Cell]bool{},
		Start:  Cell{Row: 0, Col: 0},
		End:    Cell{Row: 0, Col: 4},
	}
}

func TestDijkstraShortestPathOnOpenCorridor(t *testing.T) {
	g := openCorridor()
	algo := Dijkstra(g)
	if algo == nil || len(algo.Steps) == 0 {
		t.Fatal("expected Dijkstra steps")
	}
	last := algo.Steps[len(algo.Steps)-1]
	if len(last.Path) != 5 {
		t.Fatalf("shortest path visits 5 cells; got %d path=%v", len(last.Path), last.Path)
	}
}

func TestBFSShortestPathOnOpenCorridor(t *testing.T) {
	g := openCorridor()
	algo := BFS(g)
	last := algo.Steps[len(algo.Steps)-1]
	if len(last.Path) != 5 {
		t.Fatalf("BFS shortest path length: got %d want 5 (path=%v)", len(last.Path), last.Path)
	}
}

func TestAStarMatchesDijkstraCostOnUniformGrid(t *testing.T) {
	g := openCorridor()
	aAlgo := AStar(g)
	dAlgo := Dijkstra(g)
	aLast := aAlgo.Steps[len(aAlgo.Steps)-1]
	dLast := dAlgo.Steps[len(dAlgo.Steps)-1]
	if len(aLast.Path) != len(dLast.Path) {
		t.Fatalf("A* path len %d != Dijkstra path len %d", len(aLast.Path), len(dLast.Path))
	}
}

package pathfinding

import "testing"

func TestDijkstraShortestPathOnOpenCorridor(t *testing.T) {
	g := &Grid{
		Width:  5,
		Height: 1,
		Walls:  map[Cell]bool{},
		Start:  Cell{Row: 0, Col: 0},
		End:    Cell{Row: 0, Col: 4},
	}

	algo := Dijkstra(g)
	if algo == nil || len(algo.Steps) == 0 {
		t.Fatal("expected Dijkstra steps")
	}
	last := algo.Steps[len(algo.Steps)-1]
	if len(last.Path) != 5 {
		t.Fatalf("shortest path visits 5 cells; got %d path=%v", len(last.Path), last.Path)
	}
}

package sorting

import (
	"slices"
	"testing"
)

func TestAlgorithmsProduceSortedFinalArray(t *testing.T) {
	original := []int{23, 11, 5, 42, 17, 8, 3, 30, 1, 19}
	tests := []struct {
		name string
		run  func([]int) *Algorithm
	}{
		{"Bubble", BubbleSort},
		{"Selection", SelectionSort},
		{"Insertion", InsertionSort},
		{"Shell", ShellSort},
		{"Quick", QuickSort},
		{"Merge", MergeSort},
		{"Heap", HeapSort},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := slices.Clone(original)
			algo := tt.run(input)
			if algo == nil || len(algo.Steps) == 0 {
				t.Fatal("expected non-empty steps")
			}
			final := algo.Steps[len(algo.Steps)-1].Array
			if !slices.IsSorted(final) {
				t.Fatalf("final state not sorted: %v", final)
			}
			if len(final) != len(original) {
				t.Fatalf("length mismatch: got %d want %d", len(final), len(original))
			}
		})
	}
}

func TestGenerateRandomArrayUniqueHeights(t *testing.T) {
	for _, size := range []int{2, 15, 50} {
		arr := GenerateRandomArray(size)
		if len(arr) != size {
			t.Fatalf("size %d: got len %d", size, len(arr))
		}
		seen := make(map[int]struct{}, size)
		for _, v := range arr {
			if v < 1 || v > 50 {
				t.Fatalf("size %d: value %d out of [1,50]", size, v)
			}
			if _, dup := seen[v]; dup {
				t.Fatalf("size %d: duplicate height %d in %v", size, v, arr)
			}
			seen[v] = struct{}{}
		}
	}
}

func TestAlgorithmsSetMoveStatLabels(t *testing.T) {
	arr := []int{3, 1, 2}
	algorithms := map[string]*Algorithm{
		"Bubble":    BubbleSort(slices.Clone(arr)),
		"Selection": SelectionSort(slices.Clone(arr)),
		"Insertion": InsertionSort(slices.Clone(arr)),
		"Shell":     ShellSort(slices.Clone(arr)),
		"Quick":     QuickSort(slices.Clone(arr)),
		"Merge":     MergeSort(slices.Clone(arr)),
		"Heap":      HeapSort(slices.Clone(arr)),
	}
	want := map[string]string{
		"Bubble": "Swp", "Selection": "Swp", "Insertion": "Shf",
		"Shell": "Swp", "Quick": "Swp", "Merge": "Wrt", "Heap": "Swp",
	}
	for name, algo := range algorithms {
		if algo.MoveStatLabel != want[name] {
			t.Fatalf("%s: MoveStatLabel %q, want %q", name, algo.MoveStatLabel, want[name])
		}
	}
}

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

package sorting

import (
	"fmt"
	"math/rand/v2"
)

// Step represents a single step in the sorting visualization
type Step struct {
	Array       []int
	Comparing   []int // Indices being compared
	Swapping    []int // Indices being swapped
	Sorted      []int // Indices that are sorted
	Pivot       int   // Pivot index for quicksort (-1 if none)
	Description string
	Comparisons int // Total comparisons so far
	Swaps       int // Total swaps so far
}

// Algorithm represents a sorting algorithm
type Algorithm struct {
	Name         string
	Description  string
	TimeComplex  string
	SpaceComplex string
	Steps        []Step
}

// GenerateRandomArray creates a shuffled array with unique values for clearer visualization
func GenerateRandomArray(size int) []int {
	// Create array with unique, evenly distributed values
	arr := make([]int, size)
	maxVal := size
	for i := range arr {
		arr[i] = i + 1
	}
	// Shuffle using Fisher-Yates
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	// Scale values to fill height better
	scale := float64(50) / float64(maxVal)
	for i := range arr {
		arr[i] = int(float64(arr[i])*scale) + 1
	}
	return arr
}

// copyArray creates a copy of an integer slice
func copyArray(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	return result
}

// BubbleSort generates steps for bubble sort visualization
func BubbleSort(arr []int) *Algorithm {
	algo := &Algorithm{
		Name:         "Bubble Sort",
		Description:  "Repeatedly swaps adjacent elements if they are in wrong order",
		TimeComplex:  "O(n²)",
		SpaceComplex: "O(1)",
		Steps:        []Step{},
	}

	data := copyArray(arr)
	n := len(data)
	sorted := []int{}
	comparisons, swaps := 0, 0

	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Pivot:       -1,
		Description: fmt.Sprintf("Starting Bubble Sort with %d elements", n),
	})

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			comparisons++
			// Comparing step
			algo.Steps = append(algo.Steps, Step{
				Array:       copyArray(data),
				Comparing:   []int{j, j + 1},
				Sorted:      copyArray(sorted),
				Pivot:       -1,
				Description: fmt.Sprintf("Compare: %d vs %d", data[j], data[j+1]),
				Comparisons: comparisons,
				Swaps:       swaps,
			})

			if data[j] > data[j+1] {
				swaps++
				swapped = true
				// Swapping step
				data[j], data[j+1] = data[j+1], data[j]
				algo.Steps = append(algo.Steps, Step{
					Array:       copyArray(data),
					Swapping:    []int{j, j + 1},
					Sorted:      copyArray(sorted),
					Pivot:       -1,
					Description: fmt.Sprintf("Swap: %d ↔ %d (out of order)", data[j+1], data[j]),
					Comparisons: comparisons,
					Swaps:       swaps,
				})
			}
		}
		sorted = append([]int{n - 1 - i}, sorted...)

		if !swapped {
			for k := 0; k < n-1-i; k++ {
				if !containsInt(sorted, k) {
					sorted = append(sorted, k)
				}
			}
			break
		}
	}

	// Final sorted state
	allSorted := make([]int, n)
	for i := 0; i < n; i++ {
		allSorted[i] = i
	}
	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Sorted:      allSorted,
		Pivot:       -1,
		Description: fmt.Sprintf("Sorted! %d comparisons, %d swaps", comparisons, swaps),
		Comparisons: comparisons,
		Swaps:       swaps,
	})

	return algo
}

// SelectionSort generates steps for selection sort visualization
func SelectionSort(arr []int) *Algorithm {
	algo := &Algorithm{
		Name:         "Selection Sort",
		Description:  "Finds minimum element and places it at the beginning",
		TimeComplex:  "O(n²)",
		SpaceComplex: "O(1)",
		Steps:        []Step{},
	}

	data := copyArray(arr)
	n := len(data)
	sorted := []int{}
	comparisons, swaps := 0, 0

	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Pivot:       -1,
		Description: fmt.Sprintf("Starting Selection Sort with %d elements", n),
	})

	for i := 0; i < n-1; i++ {
		minIdx := i

		algo.Steps = append(algo.Steps, Step{
			Array:       copyArray(data),
			Pivot:       minIdx,
			Sorted:      copyArray(sorted),
			Description: fmt.Sprintf("Finding minimum starting from index %d (value: %d)", i, data[i]),
			Comparisons: comparisons,
			Swaps:       swaps,
		})

		for j := i + 1; j < n; j++ {
			comparisons++
			isNewMin := data[j] < data[minIdx]
			algo.Steps = append(algo.Steps, Step{
				Array:       copyArray(data),
				Comparing:   []int{minIdx, j},
				Sorted:      copyArray(sorted),
				Pivot:       minIdx,
				Description: fmt.Sprintf("Compare: %d vs %d (current min: %d)", data[minIdx], data[j], data[minIdx]),
				Comparisons: comparisons,
				Swaps:       swaps,
			})

			if isNewMin {
				minIdx = j
				algo.Steps = append(algo.Steps, Step{
					Array:       copyArray(data),
					Pivot:       minIdx,
					Sorted:      copyArray(sorted),
					Description: fmt.Sprintf("New minimum found: %d at index %d", data[minIdx], minIdx),
					Comparisons: comparisons,
					Swaps:       swaps,
				})
			}
		}

		if minIdx != i {
			swaps++
			data[i], data[minIdx] = data[minIdx], data[i]
			algo.Steps = append(algo.Steps, Step{
				Array:       copyArray(data),
				Swapping:    []int{i, minIdx},
				Sorted:      copyArray(sorted),
				Pivot:       -1,
				Description: fmt.Sprintf("Swap: place minimum %d at position %d", data[i], i),
				Comparisons: comparisons,
				Swaps:       swaps,
			})
		}

		sorted = append(sorted, i)
	}

	sorted = append(sorted, n-1)
	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Sorted:      sorted,
		Pivot:       -1,
		Description: fmt.Sprintf("Sorted! %d comparisons, %d swaps", comparisons, swaps),
		Comparisons: comparisons,
		Swaps:       swaps,
	})

	return algo
}

// InsertionSort generates steps for insertion sort visualization
func InsertionSort(arr []int) *Algorithm {
	algo := &Algorithm{
		Name:         "Insertion Sort",
		Description:  "Builds sorted array one element at a time by inserting each element into its correct position",
		TimeComplex:  "O(n²)",
		SpaceComplex: "O(1)",
		Steps:        []Step{},
	}

	data := copyArray(arr)
	n := len(data)
	comparisons, swaps := 0, 0

	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Sorted:      []int{0},
		Pivot:       -1,
		Description: fmt.Sprintf("Starting Insertion Sort with %d elements. First element is trivially sorted.", n),
	})

	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1

		// Show sorted portion so far
		sortedSoFar := make([]int, i)
		for k := 0; k < i; k++ {
			sortedSoFar[k] = k
		}

		algo.Steps = append(algo.Steps, Step{
			Array:       copyArray(data),
			Pivot:       i,
			Sorted:      sortedSoFar,
			Description: fmt.Sprintf("Insert %d into sorted portion [0..%d]", key, i-1),
			Comparisons: comparisons,
			Swaps:       swaps,
		})

		shifted := false
		for j >= 0 && data[j] > key {
			comparisons++
			swaps++
			algo.Steps = append(algo.Steps, Step{
				Array:       copyArray(data),
				Comparing:   []int{j, j + 1},
				Pivot:       i,
				Sorted:      sortedSoFar,
				Description: fmt.Sprintf("Shift: %d > %d, move %d right", data[j], key, data[j]),
				Comparisons: comparisons,
				Swaps:       swaps,
			})
			data[j+1] = data[j]
			j--
			shifted = true

			// Show intermediate state
			algo.Steps = append(algo.Steps, Step{
				Array:       copyArray(data),
				Swapping:    []int{j + 1, j + 2},
				Pivot:       -1,
				Sorted:      sortedSoFar,
				Description: "Element shifted, checking next position",
				Comparisons: comparisons,
				Swaps:       swaps,
			})
		}

		if j >= 0 {
			comparisons++ // Count the final comparison that ended the loop
		}

		data[j+1] = key

		// Update sorted portion to include new element
		newSorted := make([]int, i+1)
		for k := 0; k <= i; k++ {
			newSorted[k] = k
		}

		desc := fmt.Sprintf("Placed %d at index %d", key, j+1)
		if !shifted {
			desc = fmt.Sprintf("%d already in correct position", key)
		}
		algo.Steps = append(algo.Steps, Step{
			Array:       copyArray(data),
			Sorted:      newSorted,
			Pivot:       -1,
			Description: desc,
			Comparisons: comparisons,
			Swaps:       swaps,
		})
	}

	allSorted := make([]int, n)
	for i := 0; i < n; i++ {
		allSorted[i] = i
	}
	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Sorted:      allSorted,
		Pivot:       -1,
		Description: fmt.Sprintf("Sorted! %d comparisons, %d shifts", comparisons, swaps),
		Comparisons: comparisons,
		Swaps:       swaps,
	})

	return algo
}

// QuickSort generates steps for quicksort visualization
func QuickSort(arr []int) *Algorithm {
	algo := &Algorithm{
		Name:         "Quick Sort",
		Description:  "Divide and conquer: partition around pivot, recursively sort subarrays",
		TimeComplex:  "O(n log n) avg",
		SpaceComplex: "O(log n)",
		Steps:        []Step{},
	}

	data := copyArray(arr)
	sorted := []int{}
	stats := &sortStats{}

	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Pivot:       -1,
		Description: fmt.Sprintf("Starting Quick Sort with %d elements", len(data)),
	})

	var quicksort func(low, high int)
	quicksort = func(low, high int) {
		if low < high {
			pivot := partitionWithStats(data, low, high, algo, &sorted, stats)
			quicksort(low, pivot-1)
			quicksort(pivot+1, high)
		} else if low == high && low >= 0 && low < len(data) {
			if !containsInt(sorted, low) {
				sorted = append(sorted, low)
			}
		}
	}

	quicksort(0, len(data)-1)

	// Mark all as sorted
	allSorted := make([]int, len(data))
	for i := range allSorted {
		allSorted[i] = i
	}
	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Sorted:      allSorted,
		Pivot:       -1,
		Description: fmt.Sprintf("Sorted! %d comparisons, %d swaps", stats.comparisons, stats.swaps),
		Comparisons: stats.comparisons,
		Swaps:       stats.swaps,
	})

	return algo
}

// ShellSort generates steps for Shell sort visualization (Sedgewick / Knuth gap sequence).
func ShellSort(arr []int) *Algorithm {
	algo := &Algorithm{
		Name:         "Shell Sort",
		Description:  "Insertion sort on interleaved subsequences with shrinking gaps — much faster than plain insertion on larger inputs",
		TimeComplex:  "O(n^1.3) typical",
		SpaceComplex: "O(1)",
		Steps:        []Step{},
	}

	data := copyArray(arr)
	n := len(data)
	comparisons, swaps := 0, 0

	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Pivot:       -1,
		Description: fmt.Sprintf("Starting Shell Sort with %d elements", n),
	})

	gap := 1
	for gap < n/3 {
		gap = 3*gap + 1
	}

	for gap >= 1 {
		algo.Steps = append(algo.Steps, Step{
			Array:       copyArray(data),
			Pivot:       -1,
			Description: fmt.Sprintf("Sorting with gap %d", gap),
			Comparisons: comparisons,
			Swaps:       swaps,
		})

		for i := gap; i < n; i++ {
			j := i
			for j >= gap {
				comparisons++
				algo.Steps = append(algo.Steps, Step{
					Array:       copyArray(data),
					Comparing:   []int{j - gap, j},
					Pivot:       i,
					Description: fmt.Sprintf("Gap %d: compare %d vs %d", gap, data[j-gap], data[j]),
					Comparisons: comparisons,
					Swaps:       swaps,
				})

				if data[j-gap] > data[j] {
					swaps++
					data[j], data[j-gap] = data[j-gap], data[j]
					algo.Steps = append(algo.Steps, Step{
						Array:       copyArray(data),
						Swapping:    []int{j - gap, j},
						Pivot:       i,
						Description: fmt.Sprintf("Swap across gap %d", gap),
						Comparisons: comparisons,
						Swaps:       swaps,
					})
					j -= gap
				} else {
					break
				}
			}
		}

		if gap == 1 {
			break
		}
		gap /= 3
	}

	allSorted := make([]int, n)
	for i := 0; i < n; i++ {
		allSorted[i] = i
	}
	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Sorted:      allSorted,
		Pivot:       -1,
		Description: fmt.Sprintf("Sorted! %d comparisons, %d swaps", comparisons, swaps),
		Comparisons: comparisons,
		Swaps:       swaps,
	})

	return algo
}

type sortStats struct {
	comparisons int
	swaps       int
}

func containsInt(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func partitionWithStats(data []int, low, high int, algo *Algorithm, sorted *[]int, stats *sortStats) int {
	pivot := data[high]

	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Pivot:       high,
		Sorted:      copyArray(*sorted),
		Description: fmt.Sprintf("Partition [%d..%d]: pivot = %d", low, high, pivot),
		Comparisons: stats.comparisons,
		Swaps:       stats.swaps,
	})

	i := low - 1

	for j := low; j < high; j++ {
		stats.comparisons++
		lessThanPivot := data[j] <= pivot
		algo.Steps = append(algo.Steps, Step{
			Array:       copyArray(data),
			Comparing:   []int{j, high},
			Pivot:       high,
			Sorted:      copyArray(*sorted),
			Description: fmt.Sprintf("Compare: %d %s %d (pivot)", data[j], cmpSymbol(lessThanPivot), pivot),
			Comparisons: stats.comparisons,
			Swaps:       stats.swaps,
		})

		if lessThanPivot {
			i++
			if i != j {
				stats.swaps++
				data[i], data[j] = data[j], data[i]
				algo.Steps = append(algo.Steps, Step{
					Array:       copyArray(data),
					Swapping:    []int{i, j},
					Pivot:       high,
					Sorted:      copyArray(*sorted),
					Description: fmt.Sprintf("Swap: move %d to left partition", data[i]),
					Comparisons: stats.comparisons,
					Swaps:       stats.swaps,
				})
			}
		}
	}

	if i+1 != high {
		stats.swaps++
	}
	data[i+1], data[high] = data[high], data[i+1]
	*sorted = append(*sorted, i+1)

	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Swapping:    []int{i + 1, high},
		Sorted:      copyArray(*sorted),
		Pivot:       -1,
		Description: fmt.Sprintf("Pivot %d in final position %d", pivot, i+1),
		Comparisons: stats.comparisons,
		Swaps:       stats.swaps,
	})

	return i + 1
}

func cmpSymbol(less bool) string {
	if less {
		return "≤"
	}
	return ">"
}

// MergeSort generates steps for merge sort visualization
func MergeSort(arr []int) *Algorithm {
	algo := &Algorithm{
		Name:         "Merge Sort",
		Description:  "Divide array in half recursively, then merge sorted halves back together",
		TimeComplex:  "O(n log n)",
		SpaceComplex: "O(n)",
		Steps:        []Step{},
	}

	data := copyArray(arr)
	stats := &sortStats{}

	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Pivot:       -1,
		Description: fmt.Sprintf("Starting Merge Sort with %d elements", len(data)),
	})

	mergeSortWithStats(data, 0, len(data)-1, algo, stats)

	allSorted := make([]int, len(data))
	for i := range allSorted {
		allSorted[i] = i
	}
	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Sorted:      allSorted,
		Pivot:       -1,
		Description: fmt.Sprintf("Sorted! %d comparisons, %d writes", stats.comparisons, stats.swaps),
		Comparisons: stats.comparisons,
		Swaps:       stats.swaps,
	})

	return algo
}

func mergeSortWithStats(data []int, left, right int, algo *Algorithm, stats *sortStats) {
	if left < right {
		mid := left + (right-left)/2

		// Highlight the range being divided
		rangeIndices := make([]int, right-left+1)
		for i := left; i <= right; i++ {
			rangeIndices[i-left] = i
		}

		algo.Steps = append(algo.Steps, Step{
			Array:       copyArray(data),
			Comparing:   rangeIndices,
			Pivot:       mid,
			Description: fmt.Sprintf("Divide [%d..%d] at mid=%d", left, right, mid),
			Comparisons: stats.comparisons,
			Swaps:       stats.swaps,
		})

		mergeSortWithStats(data, left, mid, algo, stats)
		mergeSortWithStats(data, mid+1, right, algo, stats)
		mergeWithStats(data, left, mid, right, algo, stats)
	}
}

func mergeWithStats(data []int, left, mid, right int, algo *Algorithm, stats *sortStats) {
	leftArr := make([]int, mid-left+1)
	rightArr := make([]int, right-mid)

	copy(leftArr, data[left:mid+1])
	copy(rightArr, data[mid+1:right+1])

	i, j, k := 0, 0, left

	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Pivot:       -1,
		Description: fmt.Sprintf("Merge: [%d..%d] + [%d..%d]", left, mid, mid+1, right),
		Comparisons: stats.comparisons,
		Swaps:       stats.swaps,
	})

	for i < len(leftArr) && j < len(rightArr) {
		stats.comparisons++
		if leftArr[i] <= rightArr[j] {
			data[k] = leftArr[i]
			algo.Steps = append(algo.Steps, Step{
				Array:       copyArray(data),
				Swapping:    []int{k},
				Comparing:   []int{left + i, mid + 1 + j},
				Pivot:       -1,
				Description: fmt.Sprintf("Pick %d from left (≤ %d)", leftArr[i], rightArr[j]),
				Comparisons: stats.comparisons,
				Swaps:       stats.swaps,
			})
			i++
		} else {
			data[k] = rightArr[j]
			algo.Steps = append(algo.Steps, Step{
				Array:       copyArray(data),
				Swapping:    []int{k},
				Comparing:   []int{left + i, mid + 1 + j},
				Pivot:       -1,
				Description: fmt.Sprintf("Pick %d from right (< %d)", rightArr[j], leftArr[i]),
				Comparisons: stats.comparisons,
				Swaps:       stats.swaps,
			})
			j++
		}
		stats.swaps++
		k++
	}

	for i < len(leftArr) {
		data[k] = leftArr[i]
		stats.swaps++
		i++
		k++
	}

	for j < len(rightArr) {
		data[k] = rightArr[j]
		stats.swaps++
		j++
		k++
	}

	// Show the merged result
	mergedRange := make([]int, right-left+1)
	for idx := left; idx <= right; idx++ {
		mergedRange[idx-left] = idx
	}
	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Sorted:      mergedRange,
		Pivot:       -1,
		Description: fmt.Sprintf("Merged [%d..%d] ✓", left, right),
		Comparisons: stats.comparisons,
		Swaps:       stats.swaps,
	})
}

// HeapSort generates steps for heap sort visualization
func HeapSort(arr []int) *Algorithm {
	algo := &Algorithm{
		Name:         "Heap Sort",
		Description:  "Build max-heap, then repeatedly extract maximum to end of array",
		TimeComplex:  "O(n log n)",
		SpaceComplex: "O(1)",
		Steps:        []Step{},
	}

	data := copyArray(arr)
	n := len(data)
	sorted := []int{}
	stats := &sortStats{}

	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Pivot:       -1,
		Description: fmt.Sprintf("Starting Heap Sort with %d elements", n),
	})

	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Pivot:       -1,
		Description: "Phase 1: Building max-heap (largest element at root)",
	})

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapifyWithStats(data, n, i, algo, sorted, stats)
	}

	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Pivot:       0,
		Description: fmt.Sprintf("Max-heap built! Root = %d (maximum)", data[0]),
		Comparisons: stats.comparisons,
		Swaps:       stats.swaps,
	})

	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Pivot:       -1,
		Description: "Phase 2: Extract max elements one by one",
		Comparisons: stats.comparisons,
		Swaps:       stats.swaps,
	})

	// Extract elements from heap
	for i := n - 1; i > 0; i-- {
		stats.swaps++
		data[0], data[i] = data[i], data[0]
		sorted = append([]int{i}, sorted...)

		algo.Steps = append(algo.Steps, Step{
			Array:       copyArray(data),
			Swapping:    []int{0, i},
			Sorted:      copyArray(sorted),
			Pivot:       -1,
			Description: fmt.Sprintf("Extract max %d → position %d", data[i], i),
			Comparisons: stats.comparisons,
			Swaps:       stats.swaps,
		})

		heapifyWithStats(data, i, 0, algo, sorted, stats)
	}

	sorted = append([]int{0}, sorted...)
	algo.Steps = append(algo.Steps, Step{
		Array:       copyArray(data),
		Sorted:      sorted,
		Pivot:       -1,
		Description: fmt.Sprintf("Sorted! %d comparisons, %d swaps", stats.comparisons, stats.swaps),
		Comparisons: stats.comparisons,
		Swaps:       stats.swaps,
	})

	return algo
}

func heapifyWithStats(data []int, n, i int, algo *Algorithm, sorted []int, stats *sortStats) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n {
		stats.comparisons++
		algo.Steps = append(algo.Steps, Step{
			Array:       copyArray(data),
			Comparing:   []int{largest, left},
			Sorted:      copyArray(sorted),
			Pivot:       i,
			Description: fmt.Sprintf("Heap: compare node %d with left child %d", data[largest], data[left]),
			Comparisons: stats.comparisons,
			Swaps:       stats.swaps,
		})
		if data[left] > data[largest] {
			largest = left
		}
	}

	if right < n {
		stats.comparisons++
		algo.Steps = append(algo.Steps, Step{
			Array:       copyArray(data),
			Comparing:   []int{largest, right},
			Sorted:      copyArray(sorted),
			Pivot:       i,
			Description: fmt.Sprintf("Heap: compare %d with right child %d", data[largest], data[right]),
			Comparisons: stats.comparisons,
			Swaps:       stats.swaps,
		})
		if data[right] > data[largest] {
			largest = right
		}
	}

	if largest != i {
		stats.swaps++
		data[i], data[largest] = data[largest], data[i]
		algo.Steps = append(algo.Steps, Step{
			Array:       copyArray(data),
			Swapping:    []int{i, largest},
			Sorted:      copyArray(sorted),
			Pivot:       -1,
			Description: fmt.Sprintf("Sift down: swap %d ↔ %d", data[largest], data[i]),
			Comparisons: stats.comparisons,
			Swaps:       stats.swaps,
		})
		heapifyWithStats(data, n, largest, algo, sorted, stats)
	}
}

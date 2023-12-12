package barpear

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestRandomPositiveIntegerSliceUpToMax_WithSeed(t *testing.T) {
	tests := []struct {
		name     string
		max      int
		seed     int64
		expected []int
	}{
		{
			name:     "Seed42",
			max:      5,
			seed:     int64(42),
			expected: []int{4, 1, 3, 0, 2, 5},
		},
		{
			name:     "seedTime",
			max:      5,
			seed:     int64(43),
			expected: []int{4, 2, 1, 0, 5, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			randomSlice := RandomPositiveIntegerSliceUpToMax(tt.max, WithSeed(tt.seed))

			// Print debug information only when running with -v flag
			if testing.Verbose() {
				fmt.Println(randomSlice)
			}

			// Check if the length of the generated slice is as expected
			if len(randomSlice) != tt.max+1 {
				t.Errorf("Unexpected length of the generated slice. Expected: %d, Got: %d", tt.max+1, len(randomSlice))
			}

			// Check if the generated slice contains all positive integers up to max
			for i := 0; i <= tt.max; i++ {
				found := false
				for _, num := range randomSlice {
					if num == i {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Generated slice is missing the expected value: %d", i)
				}
			}

			// Check if the generated slice is shuffled by comparing with the expected slice
			if !reflect.DeepEqual(randomSlice, tt.expected) {
				t.Errorf("Generated slice is not as expected. Expected: %v, Got: %v", tt.expected, randomSlice)
			}
		})
	}
}

func TestRandomPositiveIntegerSliceUpToMax_UniqueOrder(t *testing.T) {
	max := 5
	seed42 := int64(42)
	seedTime := time.Now().UnixNano()

	// Generate slices with different seeds
	slice42 := RandomPositiveIntegerSliceUpToMax(max, WithSeed(seed42))
	sliceSeededFromTime := RandomPositiveIntegerSliceUpToMax(max, WithSeed(seedTime))

	if testing.Verbose() {
		fmt.Println("   Slice with seed42:", slice42)
		fmt.Println("Slice with time seed:", sliceSeededFromTime)
	}

	// Check if the generated slices have unique orders
	if reflect.DeepEqual(slice42, sliceSeededFromTime) {
		t.Errorf("Generated slices have the same order for different seeds. Seed42: %v, seedTime: %v", slice42, sliceSeededFromTime)
	}
}

func TestRandomPositiveIntegerSliceUpToMax_RandomOrder(t *testing.T) {
	max := 5

	// Generate a slice with the current time's UnixNano as the seed
	randomSlice := RandomPositiveIntegerSliceUpToMax(max)

	// Check if the generated slice is not in ascending or descending order
	isAscending := isSorted(randomSlice, true)
	isDescending := isSorted(randomSlice, false)

	if isAscending || isDescending {
		t.Errorf("Generated slice is in ascending or descending order. Slice: %v", randomSlice)
	}
}

// isSorted checks if the given slice is sorted in ascending or descending order.
// If ascending is true, it checks for ascending order; otherwise, it checks for descending order.
func isSorted(slice []int, ascending bool) bool {
	for i := 1; i < len(slice); i++ {
		if (ascending && slice[i-1] > slice[i]) || (!ascending && slice[i-1] < slice[i]) {
			return false
		}
	}
	return true
}

package merge_sort

import "testing"

func TestMergeSort(t *testing.T) {

	tests := []struct {
		name string
		arr  []int
	}{
		{
			name: "test-1",
			arr:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			name: "test-2",
			arr:  []int{2, 3, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Before merge sort", tt.arr)
			merge_sort(tt.arr)
			t.Log("After merge sort", tt.arr)
		})
	}

}

package selection_sort

import "testing"

func TestSelectionSort(t *testing.T) {

	tests := []struct {
		name string
		arr  []int
	}{
		{
			name: "test-1",
			arr:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "test-2",
			arr:  []int{2, 3, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Before select sort", tt.arr)
			selection_sort(tt.arr)
			t.Log("After select sort", tt.arr)
		})
	}

}

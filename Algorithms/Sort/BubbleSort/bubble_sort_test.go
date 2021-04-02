package bubble_sort

import "testing"

func TestSortBubble(t *testing.T) {

	tests := []struct {
		name string
		a    []int
	}{
		{
			name: "test-1",
			a:    []int{3, 2, 1, 4},
		},
		{
			name: "test-2",
			a:    []int{0, -10, 0, -30, -50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Before sorting:%_v", tt.a)
			bubble_sort(tt.a)
			t.Log("After sorting:%_v", tt.a)
		})
	}

}

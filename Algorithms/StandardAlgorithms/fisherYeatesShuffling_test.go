package FisherYatesShuffling

import (
	"testing"
)

func TestFisherYatesShuffling(t *testing.T) {

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
			a:    []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "test-3",
			a:    []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Before sorting:%_v", tt.a)
			FisherYatesShuffling(tt.a)
			t.Log("After sorting:%_v", tt.a)
		})
	}

}

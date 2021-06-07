package find

import "testing"

func TestPeakFinder(t *testing.T) {
	tests := []struct {
		name string
		a    []int
	}{
		{
			name: "test-1",
			a:    []int{1, 3, 20, 2, 1},
		},
		{
			name: "test-2",
			a:    []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Array:%_v", tt.a)
			t.Log("Peak Value:%_f", Peak_finder(tt.a))
		})
	}

}

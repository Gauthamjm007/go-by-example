package selection_sort

import "fmt"

// selection sort searches for the min value and swaps it with the next larger value , it runs !n-1 times, it runs as same as bubble sort but  it swaps a little lesser

func selection_sort(a []int) {

	timesRun := 0
	arrLength := len(a)
	for i := 0; i < arrLength; i++ {
		minIndex := i
		for j := i + 1; j < arrLength; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
			timesRun++
		}
		if minIndex != i {
			a[i], a[minIndex] = a[minIndex], a[i]
		}

	}
	fmt.Println(timesRun, "timesRun")
}

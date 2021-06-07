package find

func Peak_Finder_Reccursion(arr []int, low int, high int, n int) int {
	mid := low + int(high-low/2)

	if (mid == 0 || arr[mid] > arr[mid-1]) && (mid == n-1 || arr[mid] > arr[mid+1]) {
		return arr[mid]
	} else if mid > 0 && arr[mid-1] > arr[mid] {
		return Peak_Finder_Reccursion(arr, low, mid-1, n)
	} else {
		return Peak_Finder_Reccursion(arr, mid+1, high, n)
	}
}

func Peak_finder_dc(arr []int) int {
	n := len(arr)
	low := 0
	high := n - 1
	return Peak_Finder_Reccursion(arr, low, high, n)
}

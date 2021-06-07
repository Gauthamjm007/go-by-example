package find

func Peak_finder(arr []int) int {
	arr_length := len(arr)
	if arr_length == 0 || arr_length == 1 {
		return 0
	}
	if arr[0] > arr[1] {
		return arr[0]
	}
	if arr[arr_length-1] > arr[arr_length-2] {
		return arr[arr_length-1]
	}

	for i := 1; i < arr_length-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			return arr[i]
		}
	}
	return 0
}

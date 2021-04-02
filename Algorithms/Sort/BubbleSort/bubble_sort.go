package bubble_sort

// bubble sort swaps two arrays at a time and runs n-1! times, where n is length of the array

func bubble_sort(a []int) {

	arrLength := len(a)

	if arrLength == 0 {
		return
	}

	for i := 0; i < arrLength; i++ {
		flag := true
		for j := 0; j < arrLength-i-1; j++ {

			if a[j] > a[j+1] {
				a[j+1], a[j] = a[j], a[j+1]
				flag = false
			}
		}
		if flag {
			break
		}
	}

}

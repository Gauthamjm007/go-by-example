package merge_sort

func merge(l, r []int) []int {
	newArr := make([]int, 0, len(r)+len(l))

	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(newArr, r...)
		}
		if len(r) == 0 {
			return append(newArr, l...)
		}
		if l[0] <= r[0] {
			newArr = append(newArr, l[0])
		} else {
			newArr = append(newArr, r[0])
		}
	}

	return newArr
}

func merge_sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	n := len(arr) / 2
	l := merge_sort(arr[:n])
	r := merge_sort(arr[n:])
	return merge(l, r)
}

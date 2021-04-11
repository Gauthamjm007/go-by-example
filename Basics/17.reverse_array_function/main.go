package main

import "fmt"

func reverse(arr []int) []int {
	reverseArr := make([]int, 0, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		reverseArr = append(reverseArr, arr[i])
	}
	return reverseArr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Array : ", arr)
	fmt.Println("Reverse Array : ", reverse(arr))
}

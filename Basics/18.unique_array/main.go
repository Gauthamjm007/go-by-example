package main

import "fmt"

func uniqueArrayInt(arr []int) []int {
	uniqueArr := make([]int, 0, len(arr))
	checkMap := make(map[int]bool)

	for _, v := range arr {
		if _, exist := checkMap[v]; !exist {
			checkMap[v] = true
			uniqueArr = append(uniqueArr, v)
		}
	}

	return uniqueArr

}

func uniqueArrayString(arr []string) []string {
	uniqueStrArr := make([]string, 0, len(arr))
	checkString := make(map[string]bool)

	for _, v := range arr {
		if _, exist := checkString[v]; !exist {
			uniqueStrArr = append(uniqueStrArr, v)
			checkString[v] = true
		}
	}

	return uniqueStrArr

}

func main() {
	newArr := []int{1, 2, 3, 4, 5, 1, 2, 4, 4, 4, 99}
	fmt.Println("Original Array : ", newArr)
	fmt.Println("Unique Array : ", uniqueArrayInt(newArr))

	newStrArr := []string{"Johnny", "Jacob", "Justin", "Jewel", "Johnny"}
	fmt.Println("Original Array : ", newStrArr)
	fmt.Println("Unique Array : ", uniqueArrayString(newStrArr))

}

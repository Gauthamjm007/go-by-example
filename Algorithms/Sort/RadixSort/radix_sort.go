package radixSort

import (
	"fmt"
	"math"
	"strconv"
)

func radixSort(intArr []int) []int {
	tmp := make([]int, len(intArr))
	copy(tmp, intArr)
	places := find_largest_num_length(tmp)
	fmt.Println(places, "places")
	for index := range make([]int, places) {

		place := int(math.Pow(float64(10), float64(index)))
		fmt.Println(place, "place")
		count := make([]int, len(tmp)+1)

		loop(place, intArr, &count)
		intArr = aux_array(place, intArr, &count)

	}
	return intArr
}

func find_largest_num_length(arr []int) int {
	biggest := arr[0]
	for _, v := range arr {
		if v > biggest {
			biggest = v
		}
	}
	return len(strconv.Itoa(biggest))
}

func loop(divisor int, intArr []int, count *[]int) {
	fmt.Println(*count, "before loop")

	for _, v := range intArr {
		rem := (v / divisor) % 10
		(*count)[rem] += 1
	}
	(*count)[0] = (*count)[0] - 1
	fmt.Println(*count, " range loop")
	for i := 1; i < len(*count); i++ {
		(*count)[i] = (*count)[i] + (*count)[i-1]
	}
	fmt.Println(*count, " after loop")
}

/*
AuxArray generate a new array for a place value divisor
The new array will be sorted according to the place value in numbers
*/

func aux_array(divisor int, intArr []int, count *[]int) []int {
	//Start from the end
	aux := make([]int, len(intArr))
	fmt.Println(aux, "before aux")
	for i := len(intArr) - 1; i >= 0; i-- {
		//find the target significant digit, if divisor is 10,
		//find the 10th place digit in the number.
		k := (intArr[i] / divisor) % 10
		//find the value corrsponding to this index in the count array
		index := (*count)[k]
		//Now in aux array, put this number at the index
		aux[index] = intArr[i]
		fmt.Println(aux, index, "index", "after aux")
		//Decrement the count array at the kth index.
		(*count)[k]--
		//fmt.Printf("Count %v -- Aux %v --- IntArr %v\n", *count, intArr, aux)
	}

	return aux
}

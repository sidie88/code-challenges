package main

import (
	"errors"
)

// func main() {
// 	fmt.Println(SortArray(0, 3, 1, 2))
// }

func SortArray(array ...int) ([]int, error) {
	indexOf0 := 0
	if array[0] == 0 {
		array[0], array[len(array)-1] = array[len(array)-1], array[0]
		indexOf0 = len(array) - 1
	} else {
		indexOf0 = findIndexByValue(array, 0)
	}

	if indexOf0 == -1 {
		return nil, errors.New("ARRAY_CONTENT_ZERO_ELEMENT")
	}

	sortedCount := 0
	i := 0
	for sortedCount < len(array)-1 && i < len(array) {
		if array[i] == 0 || array[i] == i {
			i++
			continue
		}
		if indexOf0 == array[i] {
			sortedCount++
			array[indexOf0], array[i] = array[i], array[indexOf0]
			indexOf0 = i
			i = 0
		} else {
			i++
		}
	}

	return array, nil

}

func findIndexByValue(array []int, val int) int {
	for i, e := range array {
		if e == val {
			return i
		}
	}
	return -1
}

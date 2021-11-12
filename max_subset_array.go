package main

import (
	"sort"
)

// func subsetAmain() {
// 	fmt.Println(subsetAFinder([]int{5, 3, 2, 4, 1, 2}))
// }

func SubsetAFinder(arr []int) []int {
	sort.Ints(arr)
	arrLen := len(arr)
	midRange := arrLen / 2

	subsetA := []int{}

	sum1 := 0
	sum2 := 0

	counter1 := 0
	counter2 := 0

	traverseDone := false
	for !traverseDone {
		for i := 0; i <= midRange; i++ {
			sum1 = sum1 + arr[i]
			counter1++
		}
		for i := midRange + 1; i < arrLen; i++ {
			sum2 = sum2 + arr[i]
			counter2++
		}
		if sum2 > sum1 && counter2 < counter1 {
			subsetA = append(subsetA, arr[midRange+1:]...)
			traverseDone = true
		}
		midRange++
	}

	return subsetA
}

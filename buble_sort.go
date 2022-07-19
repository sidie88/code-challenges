package main

import "fmt"

func test() {
	array := []int{20, 35, -15, 7, 55, 1, -22}
	bubbleSort(array)
	fmt.Println(array)
}

func bubbleSort(array []int) {
	arrayLength := len(array)
	for lastUnsortedIndex := arrayLength - 1; lastUnsortedIndex > 0; lastUnsortedIndex-- {
		for i := 0; i < lastUnsortedIndex; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
			}
		}
	}

}

package main

func SelectionSort(n []int) []int {
	length := len(n)
	for i := 0; i < length; i++ {
		min := i
		temp := n[i]
		for j := i + 1; j < length; j++ {
			if n[j] < n[min] {
				min = j
			}
		}
		n[i], n[min] = n[min], temp
	}

	return n
}

// func main() {
// 	arr := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
// 	fmt.Println(SelectionSort(arr))
// }

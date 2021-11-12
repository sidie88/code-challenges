package main

var sumByIndex map[int]int = map[int]int{}

// func main() {
// 	fmt.Println(splitSum([]int{1, 2, 3, 4, 2}))
// }

func SplitSum(ar []int) bool {
	midIndex := (len(ar) - 1) / 2
	leftSum := 0
	rightSum := 0
	totalSum := sum(ar, len(ar)-1)

	for midIndex < len(ar) {
		leftSum = sum(ar, midIndex)
		rightSum = totalSum - leftSum
		if leftSum == rightSum {
			return true
		}
		if leftSum > rightSum {
			midIndex = midIndex - 1
		} else {
			midIndex = midIndex + 1
		}
	}

	return false

}

func sum(ar []int, endIndex int) int {
	if val, ok := sumByIndex[endIndex]; ok {
		return val
	}

	total := 0
	for i := 0; i <= endIndex; i++ {
		total = total + ar[i]
		if _, ok := sumByIndex[endIndex]; !ok {
			sumByIndex[i] = total
		}

	}
	return total
}

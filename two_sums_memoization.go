package main

func TwoSumMemoizations(nums []int, target int) []int {

	length := len(nums)
	maps := make(map[int]int)
	for i := 0; i < length; i++ {
		curr := nums[i]
		x := target - curr
		if _, ok := maps[x]; ok {
			return []int{maps[x], i}
		}
		maps[curr] = i
	}
	return []int{}
}

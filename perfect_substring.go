package main

// func perfectSubstringMain() {

// 	fmt.Println(perfectSubstring("1020122", 2))

// }

func PerfectSubstring(s string, k int) int {
	counter := 0
	marker := make(map[string]int)

	for i := 0; i < len(s); i++ {
		if val, ok := marker[string(s[i])]; ok {
			marker[string(s[i])] = val + 1
		} else {
			marker[string(s[i])] = 1
		}
		if isPerfect(marker, k) {
			counter++
			marker = make(map[string]int)
			i = i - 1
		}
	}

	return counter
}

func isPerfect(marker map[string]int, k int) bool {
	for _, v := range marker {
		if v < k {
			return false
		}
	}
	return true
}

package main

// func slowest() {
// 	keyTimes := [][]int{
// 		{0, 2},
// 		{1, 3},
// 		{0, 7},
// 	}
// 	fmt.Println("slowest key : ", slowesKey(keyTimes))
// }

func SlowesKey(keyTimes [][]int) string {
	key := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p",
		"q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	slow := keyTimes[0][1]
	slowKey := keyTimes[0][0]
	lenKey := len(keyTimes) - 1

	for i := 1; i <= lenKey; i++ {
		if keyTimes[i][1]-keyTimes[i-1][1] > slow {
			slow = keyTimes[i][1] - keyTimes[i-1][1]
			slowKey = keyTimes[i][0]
		}
	}
	return key[slowKey]

}

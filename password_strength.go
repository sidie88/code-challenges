package main

import "fmt"

/*
 * Complete the 'findPasswordStrength' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts STRING password as parameter.
 */

// func main() {
// 	strength := findPasswordStrength("test")
// 	fmt.Println("password strength: ", strength)
// }

func findPasswordStrength(password string) int64 {
	// Write your code here
	counter := int64(0)
	combination := createSubStringCombination(password)
	for _, v := range combination {
		counter = counter + count(v)
	}
	return counter
}

func createSubStringCombination(password string) []string {
	result := make([]string, 1)
	pswdLen := len(password)
	if pswdLen == 0 {
		return result
	}
	for gap := 1; gap <= pswdLen; gap++ {
		for i := 0; i < pswdLen; i++ {
			if gap+i > pswdLen {
				break
			}
			pswdChunk := string(password[i : gap+i])
			fmt.Println(string(pswdChunk))
			result = append(result, pswdChunk)
		}
	}
	return result
}

func count(passwordChunk string) int64 {
	fmt.Println(passwordChunk)

	var counterMap = map[uint8]int64{}
	counter := int64(0)
	for i := 0; i < len(passwordChunk); i++ {
		currentChar := passwordChunk[i]
		if _, ok := counterMap[currentChar]; !ok {
			counterMap[currentChar] = 1
			counter++
		}
	}
	return counter
}

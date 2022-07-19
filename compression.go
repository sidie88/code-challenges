package main

import (
	"strconv"
	"strings"
)

//func main() {
//	fmt.Println(compressedString("abaabbbc"))
//}

func compressedString(message string) string {
	// Write your code here
	length := len(message)
	if length == 0 {
		return ""
	}
	prevChar := message[0]
	var result strings.Builder
	result.WriteByte(prevChar)
	counter := 1

	for i := 1; i < length; i++ {
		currentChar := message[i]
		if prevChar == currentChar {
			counter++
			continue
		} else {
			if counter > 1 {
				result.WriteByte(currentChar)
				result.WriteString(strconv.Itoa(counter))
			} else {
				result.WriteByte(currentChar)
			}
			prevChar = currentChar
			counter = 1
		}

	}

	return result.String()

}

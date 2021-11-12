package main

import "fmt"

// func main() {
// 	x := count(2437, 875)
// 	fmt.Println("result", x)
// }

func Count(a int, b int) int {
	x := a
	y := b
	for x == y {
		if x > y {
			x = x - y
		}
		if x < y {
			y = y - x
		}
	}
	fmt.Println("x: ", x, "y: ", y)
	return x
}

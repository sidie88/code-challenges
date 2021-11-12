package main

// import "fmt"

var m map[int]int = map[int]int{}

// func main() {

// 	fmt.Println(fib(50))
// }

func Fib(n int) int {

	if val, ok := m[n]; ok {
		return val
	}
	if n <= 2 {
		return 1
	}

	m[n] = Fib(n-1) + Fib(n-2)
	return m[n]
}

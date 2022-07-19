package main

import "fmt"

type coordinate struct {
	x int
	y int
}

func testRotten() {
	orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})
}

func orangesRotting(grid [][]int) int {
	// cell := findRotten(grid)

	maps := make(map[coordinate]int)
	for i, y := range grid {
		for j, x := range y {
			cn := coordinate{
				x: i,
				y: j,
			}
			maps[cn] = x
			fmt.Println()
		}
	}
	count := 0
	// for {
	// 	// right
	// 	r := coordinate{
	// 		x: cell[0] + 1,
	// 		y: cell[1],
	// 	}
	// 	// left
	// 	l := coordinate{
	// 		x: cell[0] - 1,
	// 		y: cell[1],
	// 	}
	// 	// up
	// 	u := coordinate{
	// 		x: cell[0],
	// 		y: cell[1] - 1,
	// 	}
	// 	// down
	// 	d := coordinate{
	// 		x: cell[0],
	// 		y: cell[1] + 1,
	// 	}

	// }
	return count
}

func findRotten(grid [][]int) []int {
	for i, y := range grid {
		for j, x := range y {
			if x == 2 {
				return []int{i, j}
			}
		}
	}
	return nil
}

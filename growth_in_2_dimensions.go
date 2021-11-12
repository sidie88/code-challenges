package main

import (
	"strconv"
	"strings"
)

// func Growth2dimesion() {
// 	upRight := []string{"1 4", "2 3", "4 1"}
// 	fmt.Println(findMaxGridValue(upRight))

// }

func FindMaxGridValue(growth []string) int {

	globalMax := 0
	grid := [][]int{{0}, {0}}
	for i := 0; i < len(growth); i++ {
		cellRanges := strings.Split(growth[i], " ")
		row, _ := strconv.Atoi(cellRanges[0])
		column, _ := strconv.Atoi(cellRanges[1])
		for j := 0; j < row; j++ {
			if len(grid) <= j {
				grid = append(grid, []int{})
			}
			rows := grid[j]
			if len(rows) < column {
				for i := len(rows) - 1; i < column-1; i++ {
					rows = append(rows, 0)
				}
				grid[j] = rows
			}
			for k := 0; k < column; k++ {
				gridVal := grid[j][k] + 1
				grid[j][k] = gridVal
				if globalMax < gridVal {
					globalMax = gridVal
				}
			}

		}
	}

	return globalMax
}

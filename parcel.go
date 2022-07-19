package main

import "fmt"

func main() {
	maps := map[int]int{
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	}

	sorted := []int{2, 3, 4, 5}
	size := len(maps)
	counter := 0
	for size > 0 {
		for _, n := range sorted {
			for k, v := range maps {
				if maps[k] == 0 {
					delete(maps, k)
					continue
				}
				maps[k] = v - maps[n]
			}
			size = len(maps)
			if size == 0 {
				break;
			}
			counter++
		}

	}
	fmt.Println(counter)
}

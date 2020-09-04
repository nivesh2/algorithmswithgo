package main

import "fmt"

func binarySearch(a []int, b int) int {
	output := 0

	low := 0
	up := len(a) - 1

	for low <= up {
		m := (low + up) / 2

		if a[m] == b {
			output = 1
			break
		}

		if b < a[m] {
			up = m - 1
		} else {
			low = m + 1
		}
	}

	return output
}

func main() {
	output := binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 15)
	fmt.Printf("choice found: %d", output)
}

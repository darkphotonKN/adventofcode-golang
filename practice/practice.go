package practice

import (
	"fmt"
	"slices"
)

func Run() {

	numbers := []int{20, 7, 10, 3, 4, 3, 14, 7, 20, 3, 15, 4}

	fmt.Println("Sorted list:", sortArrayEvenAndOdd(numbers))
}

// Sort Function - Sort based on function provided
func sortArrayEvenAndOdd(list []int) []int {
	slices.SortFunc(list, func(a, b int) int {
		// can use a or b, doesn't matter
		if a%2 == 0 {
			return 1
		} else {
			return -1
		}
	})

	return list
}

package search

import "math"

func BinarySearch(numbers [] int, target int): bool {
	// find numbers by halfing and checking the center index
	midIndex := math.Floor(len(numbers) / 2)

}


func LinearSearch(numbers []int, target int): bool {

	for _, n := range numbers {
		if(n == target) {
			return true
		}
	}

	return false
}



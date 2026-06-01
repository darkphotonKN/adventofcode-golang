package practice

import (
	"fmt"
	"slices"
)

type User struct {
	name string
	age  int
}

func Run() {
	res := ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Printf("\nThree Sum Result: %+v\n\n", res)

	numbers := []int{20, 7, 10, 3, 4, 3, 14, 7, 20, 3, 15, 4}

	fmt.Println("Currying Review - Sorted list:", sortArrayEvenAndOdd(numbers))

	// updating pointer of slices
	users := []User{
		{
			name: "bob",
			age:  12,
		},
		{
			name: "John",
			age:  33,
		},
	}
	fmt.Printf("\n\nusers BEFORE: %+v\n\n", users)

	// find john
	var foundUser *User

	for _, user := range users {
		if user.name == "John" {
			foundUser = &user
		}
	}

	// update john's age
	foundUser.age = 69

	fmt.Printf("\n\nusers AFTER: %+v\n\n", users)
	currying()
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

// Given an array of strings strs, group the anagrams together. Return the groups in any order.
// Input:  ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
func GroupAnagrams(words []string) [][]string {
	// have a strategy to sort the string to create a unique "key" to know if they match
	// iterate through anagrams once and check if the sorted unique keys match
	anagramGroups := make(map[string][]string)

	for _, word := range words {
		// word sorting is O(log n) time complexity however
		sortedKey := word
		slices.Sort([]byte(sortedKey))

		// add it to the list that matches
		anagramGroups[sortedKey] = append(anagramGroups[sortedKey], word)
	}

	// convert back to slice of string slices
	res := make([][]string, 0, len(anagramGroups))

	for _, anagrams := range anagramGroups {
		res = append(res, anagrams)
	}
	return res
}

// Given an integer array nums, return all unique triplets [nums[i], nums[j], nums[k]] where the three indices are distinct and the values sum to 0. The result must contain no duplicate triplets.
// Input:  [-1, 0, 1, 2, -1, -4]
// Output: [[-1, -1, 2], [-1, 0, 1]]
//
// Input:  [0, 0, 0]
// Output: [[0, 0, 0]]
//
// Input:  [0, 1, 1]
// Output: []
// Constraints: 3 ≤ len(nums) ≤ 3000, -10^5 ≤ nums[i] ≤ 10^5.

func ThreeSum(nums []int) [][]int {
	res := make([][]int, 0)

	// sort array O(log n)
	slices.Sort(nums)

	// loop through ascending order
	triplet := make([]int, 0, 3) // initalize first triplet
	leftIndex := 1
	rightIndex := len(nums) - 1

	for curr := 0; curr < len(nums); curr++ {
		// pointer going forward
		leftNum := nums[leftIndex]
		rightNum := nums[rightIndex]
		currNum := nums[curr]

		// add all 3 and check for sum
		sum := currNum + leftNum + rightNum
		fmt.Printf("\nsum this iteration: %d\n\n", sum)
		if sum == 0 {
			// add to triplet
			triplet = append(triplet, currNum)
			triplet = append(triplet, leftNum)
			triplet = append(triplet, rightNum)

			// current triplet finished, add to result
			res = append(res, triplet)

			// initialize a fresh one
			triplet = make([]int, 0, 3)
		}

		// check the direction of the miss, is it lower than 0 or higher than 0, to change the direction we're scanning for a match
		if sum > 0 {
			if leftIndex > 0 {
				leftIndex--
			}
			if rightIndex > 0 {
				rightIndex--
			}
			continue
		}

		// else its smaller than 0, so increment to attempt to get closer to 0
		if leftIndex < len(nums)-1 {
			leftIndex++
		}

		if rightIndex < len(nums)-1 {
			rightIndex++
		}
	}

	return res
}

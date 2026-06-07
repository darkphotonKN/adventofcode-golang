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
	res := ThreeSum([]int{0, 0, 0, 0})
	fmt.Printf("\nThree Sum Result: %+v\n\n", res)
	res2 := ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Printf("\nThree Sum Result 2: %+v\n\n", res2)

	res3 := ThreeSum([]int{-1, -1, -1, 0, 0, 0, 1, 1, 1})
	fmt.Printf("\nThree Sum Result 3: %+v\n\n", res3)

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

// Given an integer array nums, return jall unique triplets [nums[i], nums[j], nums[k]] where the three indices are distinct and the values sum to 0. The result must contain no duplicate triplets.
// Input:  [-1, 0, 1, 2, -1, -4]
// Output: [[-1, -1, 2], [-1, 0, 1]]
//
// Input:  [0, 0, 0]
// Output: [[0, 0, 0]]
//
// Input:  [0, 1, 1]
// Output: []
// Constraints: 3 ≤ len(nums) ≤ 3000, -10^5 ≤ nums[i] ≤ 10^5.

// this two pointer solution space complexity is O(1)
func ThreeSum(nums []int) [][]int {

	res := make([][]int, 0)

	// sort with O(log n) cost but allows for the solution
	slices.Sort(nums)

	// two babies, one pointer from the left most and one on the right most converging

	// outer loop to find the target on each iteration
	// only check up to -2 as we'll nevber need to check the final two or one number
	// final 3 numbers would have been the final 3 numbers to check
	for i := 0; i < len(nums)-2; i++ {

		// every iteration we reset these indexes
		// we never check whats behind us, left most index and right most index converge from i ONWARDS.
		// anything from before would have been tried already, as they were all i's that matched with all other combinations.
		leftBabyIndex := i + 1          // only from i onward, explained above
		rightBabyIndex := len(nums) - 1 // always start from the final index so we get all combinations

		// skip if same as previous i, as it would net the same findings
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// also: make sure that left baby and right baby don't overlap
		for leftBabyIndex < rightBabyIndex {

			// calculate to see if triplet's sum meets target
			sum := nums[leftBabyIndex] + nums[rightBabyIndex] + nums[i]

			if sum == 0 {
				// found, save in result
				triplet := []int{nums[leftBabyIndex], nums[rightBabyIndex], nums[i]}
				res = append(res, triplet)

				// move both pointers to keep finding more combinations
				leftBabyIndex++
				rightBabyIndex--

				// but skip the same numbers if they're the same after incrementing
				for leftBabyIndex < rightBabyIndex && nums[leftBabyIndex] == nums[leftBabyIndex-1] {
					leftBabyIndex++
				}

				for leftBabyIndex < rightBabyIndex && nums[rightBabyIndex] == nums[rightBabyIndex+1] {
					rightBabyIndex--
				}
				continue
			}

			// if sum not the target, we move only the pointer that makes sense and try again

			// sum too high remove the right index down
			if sum > 0 && rightBabyIndex > leftBabyIndex {
				rightBabyIndex--
			}

			// sum too low remove the right index down
			if sum < 0 && leftBabyIndex < rightBabyIndex {
				leftBabyIndex++
			}
		}
	}

	return res
}

//

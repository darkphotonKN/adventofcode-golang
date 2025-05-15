package questions

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/

func (c *LeetCode) twoSumNaive(nums []int, target int) []int {
	// make a hashmap with a map
	// key will be the value, value will be the index
	for outerIndex, outerNum := range nums {
		for innerIndex, innerNum := range nums {
			// skip the same index
			if outerIndex == innerIndex {
				continue
			}

			if outerNum+innerNum == target {

				return []int{outerIndex, innerIndex}
			}
		}
	}

	return nil
}

func (c *LeetCode) twoSum(nums []int, target int) []int {
	// hashmap that represents numbers seen
	seen := make(map[int]int)

	// do one pass and find the different with target, check if its in the hashmap
	for index, num := range nums {
		diff := target - num
		_, exists := seen[diff]
		// we found our match
		if exists {
			return []int{index, seen[diff]}
		}
		// else we just store the number
		seen[num] = index
	}
	return nil
}

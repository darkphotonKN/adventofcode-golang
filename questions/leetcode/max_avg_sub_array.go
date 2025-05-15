package questions

import "fmt"

/*
You are given an integer array nums consisting of n elements, and an integer k.
Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.
Example 1:

Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
Example 2:

Input: nums = [5], k = 1
Output: 5.00000

Constraints:

n == nums.length
1 <= k <= n <= 105
-104 <= nums[i] <= 104
*/

func (c *LeetCode) findMaxAverage(nums []int, k int) float64 {
	// find first average
	var maxAverage float64

	sum := 0.0

	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}

	maxAverage = sum / float64(k)

	fmt.Printf("\nstarting average %f\n\n", maxAverage)

	// shifting window to find the largest
	for i := 1; i < len(nums)-k+1; i++ {

		// slide the window based on the items that slide in and slide out
		sum -= float64(nums[i-1])
		sum += float64(nums[i+k-1])

		// compare maxAverage with current iteration average
		currAvg := (sum / float64(k))
		fmt.Printf("\ncurrent average %f\n\n", currAvg)
		if currAvg > maxAverage {
			maxAverage = currAvg
		}
	}

	return maxAverage
}

func (c *LeetCode) findMaxAverageNaive(nums []int, k int) float64 {
	// find first average
	var maxAverage float64

	sum := 0.0
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}
	maxAverage = sum / float64(k)

	fmt.Printf("\nstarting average %f\n\n", maxAverage)

	// shifting window to find the largest
	for i := 1; i < len(nums)-k+1; i++ {
		sum = 0.0
		for j := i; j < k+i; j++ {
			sum += float64(nums[j])
		}
		// compare maxAverage with current iteration average
		currAvg := (sum / float64(k))
		fmt.Printf("\ncurrent average %f\n\n", currAvg)
		if currAvg > maxAverage {
			maxAverage = currAvg
		}
	}

	return maxAverage
}

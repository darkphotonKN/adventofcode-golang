package slidingwindow

import (
	"fmt"
	"math"
)

// Core idea: maintain a contiguous range [left, right] over the input. Both pointers move same direction (left to right). As you walk right forward, the window grows; when an invariant breaks, you move left forward to shrink until the invariant holds again. You incrementally update some aggregate (running sum, count, frequency map) so each character is touched O(1) per step.That's the whole pattern. The trick is choosing the right invariant

// The two flavors
// There are two distinct shapes, and recognizing which one you're in matters:

// Fixed-size window. The window size is given upfront (find max sum of any subarray of length k). left and right move together — both shift by 1 each step. The window always has the same size.
// "compute initial window of size k, then slide one step at a time"

// Variable-size window. The window grows and shrinks based on a condition. right moves forward; left only catches up when an invariant is violated. The size varies.
// "grow right until invariant breaks → shrink left until invariant restored"
// Variable-size is the more common interview flavor. Fixed-size is simpler and appears as easies (Maximum Average Subarray I).

// The recognition triggers
// A problem is sliding window when:
//
// Contiguous subarray or substring with some property.
// "Longest," "smallest," "max," "min" + "subarray/substring with property X."
// "At most K distinct," "exactly K," "with sum equals K."
// "Find a subarray that..." (when "contiguous" is implicit).

// Anti-pattern: if the problem says "subsequence" (not contiguous), it's NOT sliding window. Subsequences allow gaps; windows don't.

// The composition with hashing (Tier 2 power move)
// Most variable-size sliding window mediums in interviews are sliding window + hashmap. The hashmap is the aggregate — frequencies of characters in the current window, or counts of distinct things.
// When you see "longest substring with property involving character counts," your brain should fire: sliding window + frequency map. That's the composition.
// Examples: Longest Substring Without Repeating Characters, Min Window Substring, Longest Substring with At Most K Distinct Characters, Permutation in String.

// SlidingWindow groups together the sliding-window-family practice problems so
// we can keep adding methods to it as we work through each one.
type SlidingWindow struct{}

// Run is the entry point for the sliding_window package. main.go calls this so
// we can swap which algorithm package is being practiced just by changing one
// line. Add new problems by writing a method on SlidingWindow (or a plain
// function) and exercising it here.
func Run() {
	s := &SlidingWindow{}

	fmt.Printf("\n--- Maximum Average Subarray I ---\n\n")
	for _, tc := range []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{5}, 1, 5},
		{[]int{0, 1, 1, 3, 3}, 4, 2.0},
		{[]int{4, 0, 2, 9, 7, 3, 8, 1, 5, 6}, 4, 6.75},
	} {
		got := s.findMaxAverage(tc.nums, tc.k)
		status := "PASS"
		if math.Abs(got-tc.want) > 1e-5 {
			status = "FAIL"
		}
		fmt.Printf("[%s] findMaxAverage(%v, %d)\n      got:  %v\n      want: %v\n\n", status, tc.nums, tc.k, got, tc.want)
	}
}

// PRACTICE PROBLEMS

// Maximum Average Subarray I (LeetCode #643, Easy).
// Given an integer array nums of length n and an integer k, find a contiguous subarray of length exactly k that has the maximum average value. Return this maximum average.
// nums = [1, 12, -5, -6, 50, 3], k = 4
// → subarrays of length 4:
//   [1, 12, -5, -6] avg = 0.5
//   [12, -5, -6, 50] avg = 12.75
//   [-5, -6, 50, 3] avg = 10.5
// → max average = 12.75
//
// nums = [5], k = 1
// → avg = 5
//
// nums = [0, 1, 1, 3, 3], k = 4
// → [0,1,1,3] avg = 1.25
//   [1,1,3,3] avg = 2.0
// → max = 2.0
// Constraints: 1 ≤ k ≤ n ≤ 10^4. Values can be negative. Answer must be within 10^-5 of correct (floating point tolerance).

func (s *SlidingWindow) findMaxAverage(nums []int, k int) float64 {
	// fixed window, fixed size, k

	// find the first subarray / window average
	var initialSum float64

	for i := 0; i < k; i++ {
		initialSum += float64(nums[i])
	}

	initialAvg := initialSum / float64(k)

	// sliding window across, comparing for the largest
	largest := initialAvg

	// 0, 1, 2, 3,  4,  5
	// 0, 8, 2, 13, 12, 17

	for i := 1; i <= len(nums)-k; i++ { // need to be inclusive, e.g. 10 length with 4 elements, final window is 6 (10-4)
		// calculate new avg
		initialSum -= float64(nums[i-1])   // subtract the one thats out of the window
		initialSum += float64(nums[i+k-1]) // add the one that comes into the window

		// calculate new avg
		newAvg := initialSum / float64(k)

		// replace if greater than current largest
		if newAvg > largest {
			largest = newAvg
		}
	}

	return largest
}

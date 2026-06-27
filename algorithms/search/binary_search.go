package search

import (
	"fmt"
)

// BinarySearch groups together the binary-search-family practice problems so we
// can keep adding methods to it as we work through each one.
type BinarySearch struct{}

// Run is the entry point for the search package. main.go calls this so we can
// swap which algorithm package is being practiced just by changing one line.
// Add new problems by writing a method on BinarySearch (or a plain function)
// and exercising it here.
func Run() {
	fmt.Printf("\n--- First Bad Version ---\n\n")
	got := firstBadVersion(5)
	fmt.Printf("firstBadVersion(5) = %d\n", got)
}

// Flavor 1
// classic search for value

// Flavor 2
// A problem is flavor 2 when all three hold:
// Asks for min or max of some parameter. "Find the minimum X such that Y." "Find the maximum Z such that W."
// Parameter has a finite, derivable range — you can identify a low and high upper bound.
// Feasibility predicate is monotonic — once the answer to "does X satisfy constraint?" flips, it stays flipped. The yes/no pattern across the parameter range is a clean N...NYY...Y or Y...YNN...N.
//
// Monotonic in flavor-2 binary search means: as the parameter increases (or decreases), the "does this satisfy the constraint?" answer changes at most once. Once it flips, it stays flipped.

// How to recognize monotonicity in any flavor-2 problem
// When you read a problem asking for min/max of some parameter, do this mental test:
//
// Pick the smallest possible value of the parameter. Does the constraint hold? (Y/N)
// Pick the largest possible value of the parameter. Does the constraint hold? (Y/N)
// If the answers are different (one Y, one N), ask: "If the constraint holds at value X, does it also hold at X+1? At X-1?"
// Capacity to Ship Packages in D Days:
//
// Parameter = ship capacity.

// Predicate = "can we ship all in ≤ D days at this capacity?"
// Monotonic? If capacity X works (ships in ≤ D days), then capacity X+1 also works (bigger ship, same or fewer days). ✅ Pattern: N N N N Y Y Y Y. Want leftmost Y (smallest capacity).
//
// Koko Eating Bananas:
//
// Parameter = eating speed k.
// Predicate = "can she finish in ≤ h hours?"
// Monotonic? If k works, k+1 also works. ✅ Same shape. Leftmost Y.
//
// Split Array Largest Sum (Shopee):
//
// Parameter = "max allowed subarray sum when splitting array into k pieces."
// Predicate = "can we split the array into ≤ k pieces such that no piece exceeds this sum?"
// Monotonic? If a larger sum works (we can split into ≤ k pieces), then any even larger sum also works (more slack). ✅ Same shape. Leftmost Y.

// PRACTICE PROBLEMS

// Fair — let's pick one where the parameter is conceptually non-numeric, to break the muscle-memory association.
// Practice problem — easy, non-arithmetic flavor 2
// First Bad Version (LeetCode #278, Easy).
// You're a product manager managing a sequence of n versions, numbered 1, 2, 3, ..., n. Each version is built on top of the previous one. At some point, one version introduced a bug, and every version after that is also bad.
// You have an API: isBadVersion(version int) bool that tells you whether a given version is bad. You want to find the first bad version — that is, the smallest version number that's bad.
// You want to minimize calls to the API (it's expensive).
// n = 5, first bad version = 4
//
// isBadVersion(1) → false
// isBadVersion(2) → false
// isBadVersion(3) → false
// isBadVersion(4) → true   ← first bad
// isBadVersion(5) → true
// Constraints: 1 ≤ n ≤ 2^31 - 1, 1 ≤ first bad ≤ n.

// Why this is flavor 2 (apply the four checks):
//
// Parameter: version number.
// Range: [1, n]. Lowest meaningful version is 1, highest is n.
// Feasibility "check": isBadVersion(v) — given to you, O(1) per call.
// Monotonicity: if version v is bad, every v+1, v+2, ... is also bad. If v is good, every v-1, v-2, ... is also good. Pattern: N N N Y Y Y Y (good→bad as version increases). Leftmost Y = first bad version. Standard skeleton.
//
// Why this one trains the right reflex:
//
// No r * r or ceil(p/k) to confuse you. The "check" is just an API call. Pure binary-search-on-the-answer.
// It looks like classic flavor 1 (searching a sequence of numbered things), but you can't see the array — isBadVersion is your only window. So it's structurally flavor 2: you binary-search the index range, using the predicate.
// API-cost framing is realistic. Real production binary searches are often about expensive lookups (database queries, network calls). This problem is essentially "find the boundary in O(log n) queries instead of O(n)."
//
// Hints to start:
//
// Use the standard leftmost-Y skeleton you already know.
// Range: left = 1, right = n.
// The "check" is literally isBadVersion(mid).
// If isBadVersion(mid) is true → mid could be the first bad version; try smaller. right = mid.
// If false → mid is good, first bad is later. left = mid + 1.
// Return left after the loop.
//
// Articulate before coding:
//
// Complexity: O(log n) API calls.
// Space: O(1).
// Why this beats linear search: 30 calls instead of 2 billion for n = 2^31.

// version:  1  2  3  4  5  6  7  8
// isBad?    N  N  N  Y  Y  Y  Y  Y
//										↑
//					answer (first Y)

// when we find a YES, where yes = confirmed bad version,
// and since its ordered from NOs monotonically to only YES's, it cant be right that
// we find the first bad version. It's between left to mid, so we update right to mid
// and keep left to dispose of the upper half
// similarly, if we find a NO at mid, it can only be the top half that contains the "first"
// bad version. Update left to mid+1 (since mid was no, it can't be mid, dispose too) and
// keep right where it is.

func isBadVersion(version int) bool {
	fmt.Printf("\nTesterson\n\n")
	return false
}

func firstBadVersion(n int) int {
	left := 1
	right := n

	// iterate until lowerbound and upperbound cross overs
	for left < right {
		// find the next mid point,
		mid := (left + right) / 2

		// check version
		isBad := isBadVersion(mid)
		if isBad {
			// keep the half we're interested in
			// mid version is already bad, it could be the first, or it could be after the first,
			// therefore everything on the right is not useful for us anymore
			right = mid
		} else {
			// otherwise if version is not bad, we haven't even found the first bad version, it must
			// be on the right, so we update the left bound up to past mid (mid was confirmed not a
			// bad version, so we excluse it)
			left = mid + 1
		}
	}

	return left
}

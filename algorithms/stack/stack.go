package stack

import (
	"fmt"
	"slices"
)

/*
(LIFO) Last in first out container
Push - add to the top
Pop - remove from the top (LAST in FIRST out)
Peek - chek the top without removing

The four problem families that scream "stack"
When you see a problem, these are the trigger shapes:
Family 1 — Matching / nesting. Brackets, tags, quotes, any structure where things open and close. The rule is always "most recent unclosed thing must match the next closer." The "most recent" part is what makes it LIFO.

Valid Parentheses — canonical.
Generate Parentheses (backtracking flavor, but uses the validity check).
Minimum Remove to Make Valid Parentheses.

Family 2 — Monotonic stack. This is the senior-flavored one — the stack holds elements that are still waiting for their answer, in strictly increasing or strictly decreasing order. When a new element arrives that "resolves" the ones waiting, you pop them. Mental model: "I'm holding things that haven't found their next-greater (or smaller) yet. The current value pays off whoever it beats."

Daily Temperatures — for each day, how many days until a warmer one.
Next Greater Element.
Largest Rectangle in Histogram (hard, but a famous monotonic stack problem).
Remove All Adjacent Duplicates in String (Shopee-tested) — pop the top when it matches incoming.

Family 3 — Expression / parsing. Evaluate or transform expressions. Push operands, and when you see an operator, pop the operands it needs, compute, push the result back.

Evaluate Reverse Polish Notation.
Basic Calculator.

Family 4 — Design with auxiliary state. Build a data structure with a "track this extra thing in O(1)" requirement. Usually you maintain a parallel stack of the auxiliary state.

Min Stack — keep a parallel stack of running minimums so GetMin() is O(1).
Max Stack.

The deeper recognition skill
The unifying trigger across all four: "the relevant item is always the most recent unresolved one." If your problem keeps asking about the most recent opener, the most recent number that hasn't found its answer yet, the most recent operand, the most recent running minimum — that's the stack signal.
Anti-pattern: when you need to access the oldest unresolved thing, that's a queue (FIFO), not a stack. Don't confuse them. Stack = recent-first, queue = oldest-first.
The senior-interview articulation move
When you reach for a stack in an interview, name which family you're invoking. "This is a matching problem so I'll use a stack." "I'll use a monotonic stack here because each element is looking for its next-greater." That framing reads as having internalized the patterns, not pattern-matched to the structure.

st = append(st, x)         // push
top := st[len(st)-1]       // peek
st = st[:len(st)-1]        // pop
len(st) == 0               // empty check
*/

// KEY INSIGHT - SKIP CREATING A REAL STACK, JUST USE A SLICE AND PRETEND ITS A STACK

// Stack groups together the stack-family practice problems so we can keep
// adding methods to it as we work through each one.
type Stack struct{}

// Run is the entry point for the stack package. main.go calls this so we can
// swap which algorithm package is being practiced just by changing one line.
// Add new problems by writing a method on Stack and exercising it here.
func Run() {
	s := &Stack{}

	fmt.Printf("\n--- Valid Parenthesis ---\n\n")
	for _, tc := range []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
		{"[", false},
		{"[(", false},
		{"]", false},
	} {
		got := s.validParenthesis(tc.input)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("[%s] validParenthesis(%q) = %v (want %v)\n", status, tc.input, got, tc.want)
	}

	fmt.Printf("\n--- Next greater element ---\n\n")
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{[]int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
	} {
		got := nextGreaterElement(tc.nums1, tc.nums2)
		status := "PASS"
		if !slices.Equal(got, tc.want) {
			status = "FAIL"
		}
		fmt.Printf("[%s] nextGreaterElement(%v, %v) = %v (want %v)\n", status, tc.nums1, tc.nums2, got, tc.want)
	}
}

// PRACTICE PROBLEMS

// Valid Parenthesis
// Given a string s of '(', ')', '{', '}', '[', ']',
// determine if every opener has a correctly-ordered, matching closer.
//
// "()"     → true
// "()[]{}" → true
// "(]"     → false
// "([)]"   → false
// "{[]}"   → true
// ""       → true
// "["      → false
// "]"      → false
//
// Family 1 — Matching / nesting: the most recent unclosed opener must match
// the next closer, so we push openers and pop when a closer arrives.
// DONT CARE ABOUT MIRROR OF THE STRING HERE, only OPEN AND CLOSE ORDER MATTERS
func (s *Stack) validParenthesis(str string) bool {
	// use stack to find the most recently seen open
	stack := make([]rune, 0)

	matchingPair := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	// iterate through the string, whenever we meet a opening brackets we need to check if the respective
	// closing one is the same
	for _, letter := range str {
		expectedOpener, isCloser := matchingPair[letter]

		if !isCloser {
			// fmt.Printf("\nadding opener to stack: %q\n\n", letter)
			// add to stack so we know the most recent closer needed
			stack = append(stack, letter)
			continue
		}

		// closer now, empty stack means no match found, closer met too early, return false
		if len(stack) == 0 {
			// fmt.Print("\nstack empty, not valid\n\n")
			return false
		}

		// check if expected opener, keyed into map by the closer, matches the stack (most recently found
		// opener)
		mostRecentOpener := stack[len(stack)-1]

		// fmt.Printf("\nchecking most recent opener from stack: %q   vs   expectedOpener: %q\n\n", mostRecentOpener, expectedOpener)

		if expectedOpener != mostRecentOpener {
			// end the check early, already know its false
			return false
		}

		// shrink after successful match
		stack = stack[:len(stack)-1]

		// fmt.Printf("\nstack after shrinking: %v\n\n", stack)
		// keep looping to check the whole string
	}

	// check if there are leftovers in stack (unmatched openers)

	return len(stack) == 0
}

// Monotonic Stack example question + practice
// Family 2 — Monotonic stack. This is the senior-flavored one — the stack holds elements that are still waiting for their answer, in strictly increasing or strictly decreasing order. When a new element arrives that "resolves" the ones waiting, you pop them. Mental model: "I'm holding things that haven't found their next-greater (or smaller) yet. The current value pays off whoever it beats."
// KEY:
// The problem that screams "monotonic stack"
// The signature problem is next greater element / next smaller element / days until warmer. Generalized:
//
// "For each element in an array, find the next (or previous) element that is larger (or smaller) than it."

// Next Greater Element I (LeetCode #496, Easy)
// You're given two distinct integer arrays nums1 and nums2, where nums1 is a subset of nums2.
// For each element of nums1, find the next greater element in nums2 (the first element to the right of it in nums2 that is strictly greater). If none exists, return -1 for that position.
// Return an array result where result[i] is the next greater element for nums1[i] in nums2.
// nums1 = [4, 1, 2]
// nums2 = [1, 3, 4, 2]
// → For 4 in nums2 (index 2), nothing to the right is greater → -1
// → For 1 in nums2 (index 0), next greater is 3
// → For 2 in nums2 (index 3), nothing to the right → -1
// Output: [-1, 3, -1]
//
// nums1 = [2, 4]
// nums2 = [1, 2, 3, 4]
// → For 2: next greater is 3
// → For 4: nothing → -1
// Output: [3, -1]
// Constraints: 1 ≤ nums1.length ≤ nums2.length ≤ 1000. All elements unique.
// Before you code — articulate to yourself:
//
// Which family? (Monotonic, obviously, but say it out loud.)
// What is the stack holding? Hint: it's elements from nums2 that are still waiting for a next-greater. What kind of monotonic order do they maintain?
// What event causes a pop, and what does the pop signify?
// How do you connect a nums2 element to its position in nums1? Hint: build a result map keyed by value (since elements are unique), then translate to nums1 at the end.
//
// One subtlety to spot before coding: you only need to walk through nums2 once. As you walk, the stack maintains "values still waiting to find someone bigger to their right." When a new value arrives, it resolves any waiting values smaller than itself. After the full walk, anything still on the stack never found a next-greater — those get -1.
// The stack will be monotonic decreasing (values waiting for a bigger value, so they sit in decreasing order from bottom to top). When a bigger value arrives, it pops all the smaller waiting values from the top.
// State the amortized O(n) when you state complexity — that's the senior move for monotonic.

// Our version of explaining the solution (insights)
// "For each value in nums1, we need to find its next-greater in nums2. Brute force walks nums2 separately per query — O(n²). The optimization is to precompute next-greater for EVERY value in nums2 in one pass, store in a map, then nums1 queries become O(1) map lookups. The one pass uses a stack to hold nums2 values still waiting for a bigger one. As we walk nums2 left-to-right, each arrival pops smaller waiters and records 'waiter → arrival' in the map. Leftovers at the end get -1. Then we walk nums1 and read the map."

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// stack strategy here to keep the smallest thing at top, largest at bottom with these rules:
	// Add the first guy in regardless.
	// Compare to current element. If smaller, found a match — pop it.
	// Keep comparing and popping the next ones.
	// THEN add the current element, which is larger than what's left.
	waitingStack := make([]int, 0)
	nextGreater := make(map[int]int)

	res := make([]int, 0)

	// iterate through nums2 ONCE to precomute all the "next greaters", storing solutions in a map
	for index, num2 := range nums2 {
		if index == 0 {
			waitingStack = append(waitingStack, num2) // first one starts in stack
			continue                                  // no next greatest check in the first round
		}

		// check if next greatest is found, and iterate till no more are found
		// peek
		for len(waitingStack) > 0 && num2 > waitingStack[len(waitingStack)-1] {
			// found, store map
			original := waitingStack[len(waitingStack)-1]
			nextGreater[original] = num2 // map the original guy waiting for his next greater to the current number

			// shrink, dont need this number on the stack anymore
			waitingStack = waitingStack[:len(waitingStack)-1]
		}

		// regardless, we need to add current number to the stack, its the first time we met this number
		waitingStack = append(waitingStack, num2)
	}

	// map over waiting stack to -1, add them to map, as remaining stack items didnt find their next greatest
	for _, stillWaiting := range waitingStack {
		nextGreater[stillWaiting] = -1
	}

	// iterate over nums1 and construct answer
	for _, num1 := range nums1 {
		mappedNum1 := nextGreater[num1]
		res = append(res, mappedNum1)
	}

	return res
}

// The Tier 1 medium

// Minimum Remove to Make Valid Parentheses (LeetCode #1249, Medium).

// Given a string s of '(', ')', and lowercase English letters, remove the minimum number of parentheses (in any positions) so that the resulting string is a valid parentheses string. Return any valid result.
// Formally, a valid parentheses string is:
//
// The empty string, or
// A string AB where A and B are valid, or
// A string (A) where A is valid.
//
// "lee(t(c)o)de)"   →   "lee(t(c)o)de"   (remove the last ')')
// "a)b(c)d"         →   "ab(c)d"         (remove the ')')
// "))(("            →   ""               (remove all four — nothing valid possible)
// "(a(b(c)d)"       →   "a(b(c)d)" or "(a(bc)d)" or "(ab(c)d)"   (remove one of the openers)
// "abc"             →   "abc"            (no parens, already valid)
// Constraints: 1 ≤ len(s) ≤ 10^5. Only '(', ')', and lowercase a-z.
// Before you code, articulate to yourself:
//
// Which family? (Matching, obviously.)
// What does the stack hold this time? Hint: not the openers themselves. There's a twist — think about what info you need to know later when removing.
// When do you know an opener is "doomed"? When do you know a closer is "doomed"?
// At the very end, what's left on the stack, and what do you do with it?

func minRemoveToMakeValid(s string) string {
	// edge cases that were claimed as "valid" by the question

	// empty string, directly return
	if s == "" {
		return s
	}

	// for tracking openers and closers
	// can key in OK = opener
	// can key in got VALUE = expected close
	matchingPair := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	stack := make([]rune, 0)

	for _, letter := range s {
		_, isOpener := matchingPair[letter]

		// if its opener, we store it in the stack and skip, waiting for closer
		if isOpener {
			stack = append(stack, letter)
			continue // skip, no expected closers
		}

		// anywhere below is dealing with checking if current element is the expected closer
		if len(stack) == 0 {
			// reached the need for checking a closer matches the most recent opener but no opener in stack, not valid
			// return false
			return ""
		}

		// peek the stack
		mostRecentOpener := stack[len(stack)-1]

		// check for expected closer of the most recent opener
		expectedCloser, ok := matchingPair[mostRecentOpener]

		if !ok {
			// return false, no match at all
			return ""
		}

		// check if the expected one matches the current letter being iterated on, if the current rune is not the expected one
		// to remove the possibility of invalid
		if expectedCloser != letter {
			// return false
			return ""
		}

		// worked, pop (shrink the stack)
		stack = stack[:len(stack)-1]
	}

	// entire string passed, return true
	return s
}

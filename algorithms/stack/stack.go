package stack

import "fmt"

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

	fmt.Println("--- Valid Parenthesis ---")
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

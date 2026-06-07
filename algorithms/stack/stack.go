package stack

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

func validParenthesis(str string) bool {

	return false
}

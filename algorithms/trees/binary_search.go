package search

type BinaryNode[T any] struct {
	value T
	left  *BinaryNode[T]
	right *BinaryNode[T]
}

// function to perform the recursion
func walk(curr *BinaryNode[int], path []int) []int {
	// consider base case, when there is BinaryNode left to recurse to
	if curr == nil {
		// finished recursing
		return path
	}

	// pre
	path = append(path, curr.value)
	return walk(curr.left, path)

	// post
	path = append(path, curr.value)
	return walk(curr.right, path)
}

// func preOrderSearch(head BinaryNode[int]) []int {
//
// }

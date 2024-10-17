package search

import "fmt"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int
}

/**
* Pre-order Traversal
* 1 Visit Node (i.e. print, get the value, etc)
* 2 Recurse Left
* 3 Recurse Right
* 4 None. No post in pre-orded traversal.
**/

// define core function
func pre_order_search(head *TreeNode) []int {
	// call recursive function
	return walk(head, []int{})
}

// define the actual function that will recurse
func walk(curr *TreeNode, path []int) []int {
	// consider base case, where we return instead of recursing
	// base cae is when curr wasn't provided - hence "nil"

	if curr == nil {
		return path
	}

	// "visit" current node
	path = append(path, curr.value)

	// recurse left, update result
	path = walk(curr.left, path)

	// recurse right, update result
	path = walk(curr.right, path)

	return path
}

func RunPreOrderBinarySearch() {
	childThreeLeft := TreeNode{
		left:  nil,
		right: nil,
		value: 18,
	}

	childThreeRight := TreeNode{
		left:  nil,
		right: nil,
		value: 21,
	}

	childTwoLeft := TreeNode{
		left:  nil,
		right: nil,
		value: 5,
	}

	childTwoRight := TreeNode{
		left:  nil,
		right: nil,
		value: 4,
	}

	childOneLeft := TreeNode{
		left:  &childTwoLeft,
		right: &childTwoRight,
		value: 23,
	}

	childOneRight := TreeNode{
		left:  &childThreeLeft,
		right: &childThreeRight,
		value: 3,
	}

	headTree := TreeNode{
		left:  &childOneLeft,
		right: &childOneRight,
		value: 17,
	}

	// run search
	traversedPath := pre_order_search(&headTree)

	fmt.Printf("Final Traversed Path: %+v\n", traversedPath)

}

package datastructure

import "fmt"

/**
* A Queue implemenation in Go.
* FIFO - first in first out.
*
* STEPS
* -- Something is Added --
* - We want to attach somethign to the Tail Side of the data structure.
*    - We point the current tail NODE's next to this new value.
*    - We point the current tail value to this new tail.
* -- Something is Removed --
* - Pop it off the front of the queue by:
*    - Assigning current head value to the "next" of the current Head.
*    - Change the next of the current head to null.
*    - Return the head node value.
*
* KEYWORDS
* Peek: Check the current head value (simply head.value).
*
**/

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
}

type Node[T any] struct {
	next  *Node[T]
	value T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		head: nil,
		tail: nil,
	}
}

// methods

func (q *Queue[T]) peek() (*T, error) {
	if q.head == nil {
		return nil, fmt.Errorf("No head value.")
	}
	return &q.head.value, nil
}

func (q *Queue[T]) enqueue(value T) {
	n := &Node[T]{
		value: value,
	}

	if q.tail == nil {
		fmt.Println("No tail exists.")
		// set both head and tail
		q.head = n
		q.tail = n
		return
	}

	// pointing the tail node's next to the new node
	q.tail.next = n
	// pointing tail to the new node
	q.tail = n
}

func (q *Queue[T]) dequeue() (*T, error) {
	if q.head == nil {
		return nil, fmt.Errorf("No head exists.")
	}

	// temp holder for current head
	currHead := q.head

	// set head to second one in queue
	q.head = currHead.next

	// check if queue is empty and reset tail too
	if q.head == nil {
		q.tail = nil
	}

	// remove the reference (if there were no garbage collection)
	// currHead.next = nil

	return &currHead.value, nil
}

func RunQueue() {

}

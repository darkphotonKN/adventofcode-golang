package practice

import (
	"fmt"
)

func add(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}


func currying() {
	add5 := add(5)
	fmt.Printf("\ncurrying result: %d\n\n", add5(7))
}

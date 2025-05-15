package questions

import "fmt"

type LeetCode struct{}

func NewLeetCode() *LeetCode {
	return &LeetCode{}
}

func (c *LeetCode) Run() {

	result := c.findMaxAverage([]int{1, 12, -5, -6, 50, 3, 10}, 4)

	fmt.Printf("\nResult was %+v\n\n", result)

}

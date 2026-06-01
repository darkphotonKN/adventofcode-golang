package questions

import "fmt"

type LeetCode struct{}

func NewLeetCode() *LeetCode {
	return &LeetCode{}
}

func (c *LeetCode) Run() {

	result := c.groupAnagrams([]string{
		"eat", "tea", "tan", "ate", "nat", "bat",
	})

	fmt.Printf("\nResult was %+v\n\n", result)

}

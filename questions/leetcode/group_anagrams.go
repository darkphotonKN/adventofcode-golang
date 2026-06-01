package questions

import (
	"sort"
	"strings"
)

/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

Example 1:
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Explanation:

There is no string in strs that can be rearranged to form "bat".
The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.
Example 2:

Input: strs = [""]

Output: [[""]]

Example 3:

Input: strs = ["a"]

Output: [["a"]]

Constraints:

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.
*/

func (c *LeetCode) groupAnagrams(strs []string) [][]string {

	anagramMap := make(map[string][]string)
	groupedStrs := make([][]string, 0)

	// iterate through strs with sliding window
	for i := 0; i < len(strs); i++ {
		// check if anagram exists

		// if it doesn't, store it with a organized key
		word := []rune(strs[i])
		sort.Slice(word, func(i, j int) bool { return word[i] < word[j] })
		key := string(word)

		// check if key exists already to store in corressponding slice
		_, exists := anagramMap[key]

		if exists {
			anagramMap[key] = append(anagramMap[key], strs[i])
		} else {
			// else add its key and create array
			anagramMap[key] = make([]string, 0)
			anagramMap[key] = append(anagramMap[key], strs[i])
		}
	}

	// convert to slice of slices
	for _, slice := range anagramMap {
		groupedStrs = append(groupedStrs, slice)
	}

	return groupedStrs
}

// create a key by bubble sorting, if we cant use build-in functions
func createKey(str string) string {
	strSlice := strings.Split(str, "")

	// bubble sort string
	for outerIndex := 0; outerIndex < len(strSlice)-1; outerIndex++ {
		for innerIndex := 0; innerIndex < len(strSlice)-1; innerIndex++ {
			if strSlice[innerIndex] > strSlice[innerIndex+1] {
				// swap
				temp := strSlice[innerIndex+1]
				strSlice[innerIndex+1] = strSlice[innerIndex]
				strSlice[innerIndex] = temp
			}
		}
	}
	return strings.Join(strSlice, "")
}

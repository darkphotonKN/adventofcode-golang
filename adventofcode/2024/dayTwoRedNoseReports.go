package adventofcode2024

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:

7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9

This example data contains six reports each containing five levels.

The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:

The levels are either all increasing or all decreasing.
Any two adjacent levels differ by at least one and at most three.
In the example above, the reports can be found safe or unsafe by checking those rules:

7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
So, in this example, 2 reports are safe.

Analyze the unusual data from the engineers. How many reports are safe?
*/

func DayTwo() {
	fmt.Println("Day Two - Red Nose Reports  - Part 1")
	file, _ := os.Open("./adventofcode/2024/data/dayTwo.txt")
	defer file.Close()
	data, _ := io.ReadAll(file)

	report := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println("report:", report)

	safeLevels := 0

	for _, levels := range report {
		levelsSlice := strings.Split(levels, " ")
		safety := validateLevels(levelsSlice)

		if safety {

			fmt.Println("Level was safe!")
			safeLevels++
		}
	}

	fmt.Printf("Safe Levels: %d\n", safeLevels)
}

func validateLevels(levels []string) bool {
	// convert string values to ints
	levelsNum := make([]int, len(levels))

	for index, str := range levels {
		number, err := strconv.Atoi(str)

		if err != nil {
			fmt.Printf("Error when converting string to int, string number was %s, error was: %s", str, err)
		}

		levelsNum[index] = number
	}

	// iterate over each level, removing it before checking safety, NOTE: but time complexity O(n^2)
	for indexToIgnore := range levelsNum {

		var testLevels []int

		// filter one level per iteration (Part 2)
		for index, level := range levelsNum {
			if index != indexToIgnore {
				testLevels = append(testLevels, level)
			}
		}

		safe := checkSafety(testLevels)

		if safe {
			return true
		}
	}

	// didnt pass all checks, return false
	return false
}

func checkSafety(levels []int) bool {
	// defaults
	allIncreasing := true
	allDecreasing := true

	var prevLevel int

	safe := true

	// iterate through entire thing first time through too

	// iterate through the filtered levels
	for index, level := range levels {
		fmt.Printf("\nIterating Levels: %v\n\n", levels)
		fmt.Printf("\nCurrent Level: %d\n\n", level)

		// skip the first level
		if index == 0 {
			prevLevel = level
			continue
		}

		// --- Check Each Rule ---
		// compare with previous and check if any of the defaults expectations are broken ---

		// -- Increasing Rule --
		if prevLevel < level {
			fmt.Printf("\nLevels Still INCREASING\n\n")
			// it must mean all decreasing is now not true, set it as so
			allDecreasing = false
		}

		// -- Decreasing Rule --
		if prevLevel > level {
			fmt.Printf("\nLevels Still DECREASING\n\n")

			// it must mean all increasing is now not true, set it as so
			allIncreasing = false
		}

		if allIncreasing {
			// check if increment is too large
			difference := math.Abs(float64(prevLevel) - float64(level))

			fmt.Println("Difference was:", difference)

			if difference == 0 || difference > 3 {
				fmt.Println("This level is unsafe")
				// unsafe
				safe = false
				break
			}
		}

		if allDecreasing {
			// check if decrement is too large
			difference := math.Abs(float64(prevLevel) - float64(level))

			fmt.Println("Difference was:", difference)

			if difference == 0 || difference > 3 {
				fmt.Println("This level is unsafe")
				// unsafe
				safe = false
				break
			}
		}

		// check if any of the rules are still holding
		if !allDecreasing && !allIncreasing {
			fmt.Println("Both all increasing and decreasing have failed, level unsafe!")
			safe = false
			break
		}

		prevLevel = level
	}

	// level safe return true as only one set of levels need to be safe
	// else continue to next set of levels
	if safe {
		return true
	}

	return false
}

/*


--- Part Two ---
The engineers are surprised by the low number of safe reports until they realize they forgot to tell you about the Problem Dampener.

The Problem Dampener is a reactor-mounted module that lets the reactor safety systems tolerate a single bad level in what would otherwise be a safe report. It's like the bad level never happened!

Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.

More of the above example's reports are now safe:

7 6 4 2 1: Safe without removing any level.
1 2 7 8 9: Unsafe regardless of which level is removed.
9 7 6 2 1: Unsafe regardless of which level is removed.
1 3 2 4 5: Safe by removing the second level, 3.
8 6 4 4 1: Safe by removing the third level, 4.
1 3 6 7 9: Safe without removing any level.
Thanks to the Problem Dampener, 4 reports are actually safe!

Update your analysis by handling situations where the Problem Dampener can remove a single level from unsafe reports. How many reports are now safe?
*/

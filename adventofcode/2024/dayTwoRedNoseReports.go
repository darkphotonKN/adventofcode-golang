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
		safety := checkSafety(levelsSlice)

		if safety {

			fmt.Println("Level was safe!")
			safeLevels++
		}
	}

	fmt.Printf("Safe Levels: %d\n", safeLevels)
}

func checkSafety(levels []string) bool {
	// convert string values to ints
	levelsNum := make([]int, len(levels))

	for index, str := range levels {
		number, err := strconv.Atoi(str)

		if err != nil {
			fmt.Printf("Error when converting string to int, string number was %s, error was: %s", str, err)
		}

		levelsNum[index] = number
	}

	// default to true
	allIncreasing := true
	allDecreasing := true

	var prevLevel int

	for index, level := range levelsNum {
		fmt.Printf("Level: %d\n", level)

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
				return false
			}
		}

		if allDecreasing {
			// check if decrement is too large
			difference := math.Abs(float64(prevLevel) - float64(level))

			fmt.Println("Difference was:", difference)

			if difference == 0 || difference > 3 {
				fmt.Println("This level is unsafe")
				// unsafe
				return false
			}
		}

		// check if any of the rules are still holding
		if !allDecreasing && !allIncreasing {
			fmt.Println("Both all increasing and decreasing have failed, level unsafe!")
			return false
		}

		prevLevel = level
	}

	// passed all checks, return true
	return true
}

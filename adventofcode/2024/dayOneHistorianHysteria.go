package adventofcode2024

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
3   4
4   3
2   5
1   3
3   9
3   3

PART 1

Within each pair, figure out how far apart the two numbers are; you'll need to add up all of those distances.
For example, if you pair up a 3 from the left list with a 7 from the right list, the distance apart is 4; if you pair up a 9 with a 3, the distance apart is 6.

In the example list above, the pairs and distances would be as follows:

The smallest number in the left list is 1, and the smallest number in the right list is 3. The distance between them is 2.
The second-smallest number in the left list is 2, and the second-smallest number in the right list is another 3. The distance between them is 1.
The third-smallest number in both lists is 3, so the distance between them is 0.
The next numbers to pair up are 3 and 4, a distance of 1.
The fifth-smallest numbers in each list are 3 and 5, a distance of 2.
Finally, the largest number in the left list is 4, while the largest number in the right list is 9; these are a distance 5 apart.
To find the total distance between the left list and the right list, add up the distances between all of the pairs you found.

In the example above, this is 2 + 1 + 0 + 1 + 2 + 5, a total distance of 11!

PART 2
*/

func DayOneHistorianHysteria() {
	fmt.Println("Day One - Historian Hysteria - Part 1")

	file, _ := os.Open("./adventofcode/2024/data/dayOne.txt")
	defer file.Close()
	data, _ := io.ReadAll(file)

	dataSlice := strings.Split(string(data), "\n")

	leftList, rightList, err := splitLists(dataSlice)

	if err != nil {
		log.Fatal("Location errored: ", err)
	}

	// sort both slices

	// sort left list from smallest to largest
	slices.Sort(leftList)

	// sort right list from smallest to largest
	slices.Sort(rightList)

	// loop and compare difference, add to find the final distance

	var totalDistance float64 = 0
	for index, leftItem := range leftList {
		totalDistance += math.Abs(float64(rightList[index]) - float64(leftItem))
		fmt.Printf("\nrightNumber %d and leftNumber %d difference was %d, totalDistance is now: %d\n", rightList[index], leftItem, int(math.Abs(float64(rightList[index]))-float64(leftItem)), int(totalDistance))
	}

	fmt.Println("totalDistance was:", int(totalDistance))

	// part two, count similar score
	similarityScore := CalculateSimilarityScore(leftList, rightList)

	fmt.Println("similarityScore", similarityScore)
}

func splitLists(data []string) ([]int, []int, error) {
	var leftList []int
	var rightList []int

	// create left and right slices
	for _, row := range data {
		locations := strings.Split(row, "   ")

		if len(locations) > 1 {

			leftLocation, err := strconv.Atoi(locations[0])

			if err != nil {
				return nil, nil, err
			}

			rightLocation, err := strconv.Atoi(locations[1])

			if err != nil {
				return nil, nil, err
			}

			leftList = append(leftList, leftLocation)
			rightList = append(rightList, rightLocation)
		}
	}

	return leftList, rightList, nil
}

func CalculateSimilarityScore(left, right []int) int {
	result := 0

	for _, leftNumber := range left {
		count := 0

		for _, rightNumber := range right {
			if leftNumber == rightNumber {
				count++
			}
		}

		// add the result at the end
		result += leftNumber * count
	}

	return result
}

/*
--- Part Two ---
Your analysis only confirmed what everyone feared: the two lists of location IDs are indeed very different.

Or are they?

The Historians can't agree on which group made the mistakes or how to read most of the Chief's handwriting, but in the commotion you notice an interesting detail: a lot of location IDs appear in both lists! Maybe the other numbers aren't location IDs at all but rather misinterpreted handwriting.

This time, you'll need to figure out exactly how often each number from the left list appears in the right list. Calculate a total similarity score by adding up each number in the left list after multiplying it by the number of times that number appears in the right list.


3   4
4   3
2   5
1   3
3   9
3   3

For these example lists, here is the process of finding the similarity score:

The first number in the left list is 3. It appears in the right list three times, so the similarity score increases by 3 * 3 = 9.
The second number in the left list is 4. It appears in the right list once, so the similarity score increases by 4 * 1 = 4.
The third number in the left list is 2. It does not appear in the right list, so the similarity score does not increase (2 * 0 = 0).
The fourth number, 1, also does not appear in the right list.
The fifth number, 3, appears in the right list three times; the similarity score increases by 9.
The last number, 3, appears in the right list three times; the similarity score again increases by 9.

So, for these example lists, the similarity score at the end of this process is 31 (9 + 4 + 0 + 0 + 9 + 9).

*/

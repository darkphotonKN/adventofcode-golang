package adventofcode2021

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type submarine struct {
	Data []int
}

func Submarine() {
	fmt.Println("Advent of Code 2021 Day 1 - Sonar Sweep")

	// read in context from data file
	file, err := os.Open("./adventofcode/2021/data/dayOne.txt")

	if err != nil {
		log.Fatal("Error occured when reading file:", err)
	}

	defer file.Close()

	// read the entire file
	data, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	// convert byte type to string
	text := string(data)
	var submarineTest submarine

	submarineTest.init(text)
	// fmt.Println("data:", submarineTest.Data)

	partOneAnswer := submarineTest.countIncreases()
	fmt.Println("partOneAnswer:", partOneAnswer)
}

func (sub *submarine) init(textInp string) {
	parsedData := sub.parseLine(textInp)
	numberData := make([]int, len(parsedData))

	// replace with int version
	for i, n := range parsedData {
		number, err := strconv.Atoi(n)
		if err != nil {
			continue
		}
		numberData[i] = number
	}

	sub.Data = numberData
}

func (sub *submarine) parseLine(s string) []string {
	return strings.Split(s, "\n")
}

func (sub *submarine) countIncreases() int {
	depthData := sub.Data
	var count int

	for i, depth := range depthData {
		// omit first index
		if i > 0 {
			// check if current depth is greater than previous depth
			if depth > depthData[i-1] {
				count++
			}
		}
	}
	return count
}

/* Description
As the submarine drops below the surface of the ocean, it automatically performs a sonar sweep of the nearby sea floor. On a small screen, the sonar sweep report (your puzzle input) appears: each line is a measurement of the sea floor depth as the sweep looks further and further away from the submarine.

For example, suppose you had the following report:

199
200
208
210
200
207
240
269
260
263
This report indicates that, scanning outward from the submarine, the sonar sweep found depths of 199, 200, 208, 210, and so on.

The first order of business is to figure out how quickly the depth increases, just so you know what you're dealing with - you never know if the keys will get carried into deeper water by an ocean current or a fish or something.

To do this, count the number of times a depth measurement increases from the previous measurement. (There is no measurement before the first measurement.) In the example above, the changes are as follows:

199 (N/A - no previous measurement)
200 (increased)
208 (increased)
210 (increased)
200 (decreased)
207 (increased)
240 (increased)
269 (increased)
260 (decreased)
263 (increased)
In this example, there are 7 measurements that are larger than the previous measurement.

How many measurements are larger than the previous measurement?



*/

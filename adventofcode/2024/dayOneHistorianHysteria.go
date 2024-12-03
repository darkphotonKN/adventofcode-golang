package adventofcode2024

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func DayOneHistorianHysteria() {
	fmt.Println("Day One - Historian Hysteria - Part 1")

	file, _ := os.Open("./adventofcode/2024/data/dayOne.txt")
	defer file.Close()
	data, _ := io.ReadAll(file)

	dataSlice := strings.Split(string(data), "\n")

	leftList, _, err := splitLists(dataSlice)

	if err != nil {
		log.Fatal("Location errored: ", err)
	}

	fmt.Println("leftList:", leftList)

}

func splitLists(data []string) ([]int, []int, error) {

	var leftList []int
	var rightList []int

	// create left and right slices
	for _, row := range data {
		locations := strings.Split(row, "   ")

		// check if empty
		if len(locations) != 0 {
			leftLocation, err := strconv.Atoi(locations[0])
			rightLocation, err := strconv.Atoi(locations[1])
			if err != nil {
				fmt.Println("Location not converted to an int, location was:", locations)
				return nil, nil, err
			}
			leftList = append(leftList, leftLocation)
			rightList = append(rightList, rightLocation)
		}
	}

	return leftList, rightList, nil
}

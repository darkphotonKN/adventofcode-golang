package adventofcode

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Trebuchet() {
	log.Println("Trebuchet")

	// read in context from data file
	file, err := os.Open("./adventofcode/data/dayOne2023trebuchet-data.txt")

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

	// incoming data
	var treb trebuchet
	treb.Data = text
	// fmt.Println(text)

	// convert string into an array of strings based on return
	treb.ParsedLine = treb.parseLine(treb.Data)

	// filter each string to digits only
	treb.FilteredDigits = treb.filterDigits(treb.ParsedLine)

	// loop and sum up the first and last values
	treb.Sum = treb.sumFirstAndLast(treb.FilteredDigits)

	fmt.Println("Updated sum:", treb.Sum)

	// sum up the first and last digits

	// 	var shortTestData = `two1nine
	// eightwothree
	// abcone2threexyz
	// xtwone3four
	// 4nineeightseven2
	// zoneight234
	// 7pqrstsixteen`
	//
	// 	parsedShortTestData := parseLine(shortTestData)
	// 	filteredShortTestData := filterDigits(parsedShortTestData)

	// fmt.Println(filteredShortTestData)

	fmt.Println("Trebuchet original parsed filtered digits index 3:", treb.FilteredDigits[3])
}

type trebuchet struct {
	Data           string
	ParsedLine     []string
	FilteredDigits []string
	Sum            int32
}

// convert a multi-line string into a string slice per new line
func (treb *trebuchet) parseLine(s string) []string {
	return strings.Split(s, "\n")
}

// filter string into only digits remaining inside string
func (treb *trebuchet) filterDigits(s []string) []string {
	// using make to preallocate memory for the correct sized string
	digitsOnly := make([]string, len(s))

	for i, textLine := range s {
		var builder strings.Builder

		// loop through string and read each rune and filter out digits with unicode
		for _, r := range textLine {
			if unicode.IsDigit(r) {
				builder.WriteRune(r)
			}
		}

		digitsOnly[i] = builder.String()
	}

	return digitsOnly
}

// convert each string digit to numbers
func (treb *trebuchet) convertDigitWordToNumber(s []string) []string {
	var numberSlice []string

	for i, line := range s {

		// swap the word digits to numbers
		numberSlice[i] = line
	}

	return s
}

// sums first and last digits of slice
func (treb *trebuchet) sumFirstAndLast(s []string) int32 {
	var sum int32 = 0

	for _, text := range s {
		// just continue if string is empty
		if text == "" {
			continue
		}
		// convert to the number type
		combinedDigit := string(text[0]) + string(text[len(text)-1])

		number, err := strconv.ParseInt(combinedDigit, 10, 32)

		if err != nil {
			continue
		}

		sum += int32(number)
	}

	return sum
}

package adventofcode2023

import (
	"challenges/tools"
	"fmt"
	"strconv"
	"strings"
)

func Trebuchet() {

	text := tools.FileReader("./adventofcode/2023/data/dayOne.txt")
	// 	text := `
	// 	two1nine
	// eightwothree
	// abcone2threexyz
	// xtwone3four
	// 4nineeightseven2
	// zoneight234
	// 7pqrstsixteen
	// 	`

	// map of string digits to actual digits
	strToNumDigits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	textArr := strings.Split(text, "\n")

	// store each extracted number and its index
	// extractedDigits := make(map[int]string) // storing its index as key and value as the number its converted to

	// storing converted array
	onlyFirstAndLastDigits := make([]string, len(textArr))

	for lineIndex, line := range textArr {
		// convert all occurences of every string digit to number
		lowUnset := true

		var lowIndex int
		var highIndex int

		var lowValue string
		var highValue string

		for digitStr, digit := range strToNumDigits {
			index := strings.Index(line, digitStr)
			lastIndex := strings.LastIndex(line, digitStr)

			// match found
			if index != -1 {

				// check if low is unset yet
				if lowUnset {
					lowIndex = index
					lowValue = digit

					// set to false to prevent subsequent loops
					lowUnset = false
				} else if index < lowIndex {

					// update lowest index
					lowIndex = index
					// and the lowest value
					lowValue = digit // store digit version of string
				}

				// do the same for largest index
				if lastIndex > highIndex {
					// update highest index
					highIndex = lastIndex
					// and the highest value
					highValue = digit
				}

				// fmt.Printf("Match found for digitStr %s, lowValue is now: %s, highValue is now: %s\n", digitStr, lowValue, highValue)
			}

			// do it for number digits too
			noIndex := strings.Index(line, digit)
			noLastIndex := strings.LastIndex(line, digit)

			// match found
			if noIndex != -1 {

				// check if low is unset yet
				if lowUnset {
					lowIndex = noIndex
					lowValue = digit

					// set to false to prevent subsequent loops
					lowUnset = false
				} else if noIndex < lowIndex {

					// update lowest index
					lowIndex = noIndex
					// and the lowest value
					lowValue = digit
				}

				// do the same for largest index
				if noLastIndex > highIndex {
					// update highest index
					highIndex = noLastIndex
					// and the highest value
					highValue = digit
				}
				// fmt.Printf("Match found for digit %s, lowValue is now: %s, highValue is now: %s\n", digit, lowValue, highValue)
			}
		}

		// store
		onlyFirstAndLastDigits[lineIndex] = lowValue + highValue
	}

	fmt.Println(onlyFirstAndLastDigits)

	// sum the numbers
	var count int
	for _, digit := range onlyFirstAndLastDigits {
		// account for numbers that came as single digits
		if len(digit) == 1 {
			digit = digit + digit
		}

		n, _ := strconv.Atoi(digit)

		count += n
	}

	fmt.Printf("total: %d\n", count)

}

// Original Attempt - beginning of go
// func Trebuchet() {
// 	log.Println("Advent of Code 2023 Day 1 - Trebuchet")
//
// 	// read in context from data file
// 	file, err := os.Open("./adventofcode/2023/data/dayOne.txt")
//
// 	if err != nil {
// 		log.Fatal("Error occured when reading file:", err)
// 	}
//
// 	defer file.Close()
//
// 	// read the entire file
// 	data, err := io.ReadAll(file)
//
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	// convert byte type to string
// 	text := string(data)
//
// 	// incoming data
// 	var treb trebuchet
// 	treb.Data = text
// 	// fmt.Println(text)
//
// 	// convert string into an array of strings based on return
// 	treb.ParsedLine = treb.parseLine(treb.Data)
//
// 	// filter each string to digits only
// 	treb.FilteredDigits = treb.filterDigits(treb.ParsedLine)
//
// 	// loop and sum up the first and last values
// 	treb.Sum = treb.sumFirstAndLast(treb.FilteredDigits)
//
// 	fmt.Println("Updated sum:", treb.Sum)
//
// 	// sum up the first and last digits
//
// 	// var shortTestData = `two1nine
// 	// eightwothree
// 	// abcone2threexyz
// 	// xtwone3four
// 	// 4nineeightseven2
// 	//  zoneight234
// 	// 7pqrstsixteen`
//
// 	var trebPartTwo trebuchet
// 	trebPartTwo.Data = text
// 	trebPartTwo.ParsedLine = trebPartTwo.parseLine(trebPartTwo.Data)
// 	digitsOnly := trebPartTwo.convertDigitWordToNumber(trebPartTwo.ParsedLine)
// 	trebPartTwo.Sum = trebPartTwo.sumFirstAndLast(digitsOnly)
//
// 	// fmt.Println("part two data:", trebPartTwo.Data)
// 	fmt.Println("trebPartTwo digitsOnly:", digitsOnly)
// 	fmt.Println("trebPartTwo sum:", trebPartTwo.Sum)
//
// }
//
// func getWordsToNums() map[string]string {
// 	wordNumsToDigits := map[string]string{
// 		"one":   "1",
// 		"two":   "2",
// 		"three": "3",
// 		"four":  "4",
// 		"five":  "5",
// 		"six":   "6",
// 		"seven": "7",
// 		"eight": "8",
// 		"nine":  "9",
// 	}
//
// 	return wordNumsToDigits
// }
//
// // Map of Word Numbers to Digit Numbers
// type trebuchet struct {
// 	Data           string
// 	ParsedLine     []string
// 	FilteredDigits []string
// 	Sum            int32
// }
//
// // convert a multi-line string into a string slice per new line
// func (treb *trebuchet) parseLine(s string) []string {
// 	return strings.Split(s, "\n")
// }
//
// // filter string into only digits remaining inside string
// func (treb *trebuchet) filterDigits(s []string) []string {
// 	// using make to preallocate memory for the correct sized string
// 	digitsOnly := make([]string, len(s))
//
// 	for i, textLine := range s {
// 		var builder strings.Builder
//
// 		// loop through string and read each rune and filter out digits with unicode
// 		for _, r := range textLine {
// 			if unicode.IsDigit(r) {
// 				builder.WriteRune(r)
// 			}
// 		}
//
// 		digitsOnly[i] = builder.String()
// 	}
//
// 	return digitsOnly
// }
//
// // convert each string digit to numbers
// func (treb *trebuchet) convertDigitWordToNumber(s []string) []string {
// 	numberSlice := make([]string, len(s))
// 	digitWords := getWordsToNums()
//
// 	// xtwone3four
// 	// x2ne3four
// 	// x2134
//
// 	// create a new string based on both digits and words to digits
//
// 	// map :
// 	// index - value
// 	// 5 - "3"
// 	//
// 	// 1 - "two"
// 	// 3 - "one"
// 	// 6 - "four"
// 	// sort them:
// 	// twoone3four => 2134 -> 24
//
// 	for i, line := range s {
// 		indexDigitMap := make(map[int8]string)
// 		fmt.Println("current interacted line:", line)
//
// 		// swap the word digits to numbers - using key as old string and value as new string
// 		for word, digit := range digitWords {
//
// 			// find if that word version of the digit exists
// 			index := strings.Index(line, word)
//
// 			if index != -1 {
// 				// found word, store index and digit
// 				indexDigitMap[int8(index)] = digit
// 			}
//
// 			// directly store any real digits as well
//
// 			for i, c := range line {
// 				// found digit
// 				if unicode.IsDigit(c) {
// 					// store digit and index
// 					indexDigitMap[int8(i)] = string(c)
// 				}
// 			}
// 		}
//
// 		fmt.Println("indexDigitMap result:", indexDigitMap)
//
// 		// create new slice based on order of indexes
// 		var fullDigitStr string
//
// 		// extract keys from the map
// 		keys := make([]int, 0, len(indexDigitMap))
// 		for k := range indexDigitMap {
// 			keys = append(keys, int(k))
// 		}
//
// 		// sort keys
// 		sort.Ints(keys) // modifies original keys
//
// 		// concatenate strings based on key order
// 		for _, key := range keys {
// 			fullDigitStr += indexDigitMap[int8(key)]
// 		}
//
// 		fmt.Println("fullDigitStr", fullDigitStr)
//
// 		// add this converted line to the corresponding index of the slice
// 		numberSlice[i] = fullDigitStr
// 	}
//
// 	return numberSlice
// }
//
// // sums first and last digits of slice
// func (treb *trebuchet) sumFirstAndLast(s []string) int32 {
// 	var sum int32 = 0
//
// 	for _, text := range s {
// 		// just continue if string is empty
// 		if text == "" {
// 			continue
// 		}
// 		// convert to the number type
// 		combinedDigit := string(text[0]) + string(text[len(text)-1])
//
// 		fmt.Println("current line to sum", text)
// 		fmt.Println()
// 		fmt.Printf("Summing %s with %s", string(text[0]), string(text[len(text)-1]))
// 		fmt.Println()
//
// 		number, err := strconv.ParseInt(combinedDigit, 10, 32)
//
// 		if err != nil {
// 			continue
// 		}
//
// 		sum += int32(number)
// 	}
//
// 	return sum
// }

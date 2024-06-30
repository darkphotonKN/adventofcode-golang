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

	// fmt.Println(text)

	// convert string into an array of strings based on return
	textSlice := parseLine(text)

	// filter each string to digits only
	digitsOnlySlice := filterDigits(textSlice)

	// loop and sum up the first and last values
	result := sumFirstAndLast(digitsOnlySlice)

	fmt.Println(result)

	// sum up the first and last digits
}

func parseLine(t string) []string {
	return strings.Split(t, "\n")
}

func filterDigits(s []string) []string {

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

func sumFirstAndLast(s []string) int32 {
	var sum int32 = 0

	for _, text := range s {
		// just continue if string is empty
		if text == "" {
			continue
		}
		// convert to the number type
		combinedDigit := string(text[0]) + string(text[len(text)-1])

		fmt.Println("combinedDigit:", combinedDigit)
		number, err := strconv.ParseInt(combinedDigit, 10, 32)
		fmt.Println("number:", number)

		if err != nil {
			continue
		}

		sum += int32(number)
		fmt.Println("sum so far:", sum)
	}

	return sum
}

package tools

import (
	"io"
	"log"
	"os"
)

// reads in the file and returns the file as a text string
func FileReader(fileLocation string) string {
	// read in context from data file
	file, err := os.Open(fileLocation)
	defer file.Close()

	if err != nil {
		log.Fatal("Error occured when reading file:", err)
	}

	// read the entire file
	data, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

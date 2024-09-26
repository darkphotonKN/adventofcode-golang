package designpatterns

import (
	"fmt"
)

/**
* Single Responsibility Principle
**/

func SingleResponsibilityPrinciple() {
	fmt.Println("Single Responsibility Principle")
	journal := NewJournal()

	journal.AddEntry("Test entry 1")
	journal.AddEntry("Test entry 2")
	journal.AddEntry("Test entry 3")
	journal.AddEntry("Test entry 4")
	_, err := journal.RemoveEntry(3)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("\nEntries: %v\n\n", journal.entries)
}

type Journal struct {
	entries []string
}

func NewJournal() *Journal {
	return &Journal{}
}

func (j *Journal) AddEntry(text string) int {
	j.entries = append(j.entries, text)

	return len(j.entries)
}

// breaks Single Responsiblity, since "Saving to a File" might be something
// that isn't specific to saving only journals
func (j *Journal) SaveToFile() {
	// some logic that writes to file
}

func (j *Journal) RemoveEntry(i uint) (string, error) {

	// check if index is even within the range of entries
	if i >= uint(len(j.entries)) {
		return "", fmt.Errorf("Index too large.")
	}

	updatedEntries := make([]string, len(j.entries)-1)
	var removedEntry string

	for index, entry := range j.entries {
		// don't add the removed entry
		if uint(index) == i {
			removedEntry = entry
			continue
		}
		updatedEntries = append(updatedEntries, entry)
	}

	j.entries = updatedEntries

	return removedEntry, nil
}

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
	// if using dependency injection method to solve SRP
	*FileManager
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

// 1. Using DI and embed to maintain Separation of Concerns
// doesn't break single resposibility principle,
// using Seperation of Concerns
type FileManager struct {
}

func NewFileManager() *FileManager {
	return &FileManager{}
}

func (f *FileManager) SaveToFile() {
	// write logic that writes to file
}

// correct usage - using depedency injection
func NewJournalFollowingSRP() *Journal {
	// init new File Manager
	fm := NewFileManager()

	return &Journal{
		FileManager: fm,
	}
}

// now this is useable but the concerns are seperated
func TestCorrectJournalFollowingSRP() {
	journal := NewJournalFollowingSRP()
	journal.AddEntry("Test entry!!")
	// available despite source code being in another module
	journal.SaveToFile()
}

// 2. Simply using a the method when needed
func SaveToFileStandalone() {
	// save to file logic
}

func (j *Journal) SomeJournalMethod() {
	// call it directly, but it remains decoupled
	SaveToFileStandalone()
}

// 3. FileManager remains a separate component with no embedding - BEST
type FileManager2 struct{}

func NewFileManager2() *FileManager {
	return &FileManager{}
}

func (f *FileManager) SaveToFile2(entries []string) {
	// logic to save entries to a file
	fmt.Println("Saving entries to file:", entries)
}

// 4. A Complete separation of conerns but DI for the OTHER module:
// FileManager responsible for saving entries to a file
type FileManager3 struct{}

func NewFileManager3() *FileManager3 {
	return &FileManager3{}
}

func (f *FileManager3) SaveToFile3(entries []string) {
	// Imagine this writes entries to a file
	fmt.Println("Saving entries to file:", entries)
}

// Journal method that saves its state via FileManager
// We inject the relation here as an argument, not from embedding in the base type.
func (j *Journal) Save(journalFileManager *FileManager3) {
	journalFileManager.SaveToFile3(j.entries)
}

// Usage example
func TestCorrectJournalFollowingSRP3() {
	journal := NewJournal()
	journal.AddEntry("First entry")
	journal.AddEntry("Second entry")

	// Injecting FileManager to save the journal
	fileManager := NewFileManager3()
	journal.Save(fileManager) // Journal just calls the FileManager to save itself
}

/**
* Dependency Injection Principle
**/

func DependencyInjectionPrinciple() {
	fmt.Println("Dependency Injection Principle")

	// breaking DIP
	relationships := Relationships{}

	john := Person{name: "John"}
	johnJr := Person{name: "John Jr."}

	michael := Person{name: "Michael"}
	jorge := Person{name: "Jorge"}

	relationships.AddParentAndChild(&john, &johnJr)
	relationships.AddParentAndChild(&michael, &jorge)

	// perform investigation with a High level module that BREAKS DIP
	research := Research{relationships: relationships}
	research.Investigate()

}

type Relationship string

const (
	Parent  Relationship = "parent"
	Child   Relationship = "child"
	Sibling Relationship = "sibling"
)

type Person struct {
	name string
}

// modelling relationship between two people
type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{
		from:         parent,
		relationship: Parent,
		to:           child,
	})
	r.relations = append(r.relations, Info{
		from:         child,
		relationship: Child,
		to:           parent,
	})
}

// BREAKING dependency inversion principle, HLM depending on LLM
type Research struct {
	// relations is tightly bound to Research
	// this would break if the LLM decides to change its implementation
	// e.g. change the slice to a Database
	relationships Relationships
}

func (r *Research) Investigate() {
	// relations is accessed in the tight bind to Research
	// e.g. this would error if LLM changes relations to anything but a slice
	for _, relation := range r.relationships.relations {
		// do stuff with relations
		fmt.Printf("\n%v is a %s of %v\n\n", relation.from.name, relation.relationship, relation.to.name)
	}
}

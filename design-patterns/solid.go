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
* Dependency Inversion Principle
**/
func DependencyInversionPrinciple() {
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
	matchingRelations, err := research.Investigate(Child)

	if err != nil {
		fmt.Println(err)
	}

	for _, relation := range matchingRelations {
		fmt.Printf("\n%v is a %s of %v\n\n", relation.from.name, relation.relationship, relation.to.name)
	}

	// CONFORMING to DIP

	// 1. perform investigation with a HLM that DOES NOT BREAK DIP - using constructor
	researchWithDIP := ResearchFollowingDIP{}
	matches, err := researchWithDIP.Investigate(relationships.relations, Parent)

	if err != nil {
		fmt.Println(err)
	}

	for _, relation := range matches {
		fmt.Printf("\n%v WITH DIP is a %s of %v\n\n", relation.from.name, relation.relationship, relation.to.name)
	}

	// 2. without using constructor
	TestDIPWithoutConstructor()
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

func (r *Research) Investigate(relationType Relationship) ([]Info, error) {
	var matchingRelations []Info
	// relations is accessed in the tight bind to Research
	// e.g. this would error if LLM changes relations to anything but a slice
	for _, relation := range r.relationships.relations {
		// do stuff with relations
		if relation.relationship == relationType {
			matchingRelations = append(matchingRelations, relation)
		}
	}

	// nothing found
	if len(matchingRelations) == 0 {
		return matchingRelations, fmt.Errorf("No matching relations")

	}

	return matchingRelations, nil
}

// MAINTAINING the principles of DI, HLM and LLM only depending on abstractions
type RelationshipBrowser interface {
	// defines the shape of input and output, not the details
	Investigate(relations []Info, relationshipType Relationship) ([]Info, error)
}

type ResearchFollowingDIP struct {
}

func NewResearchFollowingDIP() RelationshipBrowser {
	return &ResearchFollowingDIP{}
}

// ONE general implentation of this, not tied to RelationshipBrowser
func (r *ResearchFollowingDIP) Investigate(relations []Info, relationType Relationship) ([]Info, error) {
	var matchingRelations []Info
	// relations is accessed in the tight bind to Research
	// e.g. this would error if LLM changes relations to anything but a slice
	for _, relation := range relations {
		// do stuff with relations
		if relation.relationship == relationType {
			matchingRelations = append(matchingRelations, relation)
		}
	}

	// nothing found
	if len(matchingRelations) == 0 {
		return matchingRelations, fmt.Errorf("No matching relations")

	}

	return matchingRelations, nil
}

// Solution without using constructors
type ResearchWithInterface struct {
	// only an interface, so no tight binding
	browser RelationshipBrowser
}

type Browser struct {
}

// func (r *ResearchWithInterface) Investigate() {
// 	relationships := Relationships{}
// 	relationships.AddParentAndChild(&Person{name: "Hiroshi"}, &Person{name: "Bob"})
// 	relationships.relations = append(relationships.relations, Info{
// 		from:         &Person{name: "Ellie"},
// 		relationship: Child,
// 		to:           &Person{name: "Joel"},
// 	})
// 	relationships.AddParentAndChild(&Person{name: "John"}, &Person{name: "John Jr"})
//
// 	r.browser.Investigate(relationships.relations, Parent)
// }

// running solution without constructor
func TestDIPWithoutConstructor() {
	fmt.Println("Test DIP without constructor")

	relationships := Relationships{}
	relationships.AddParentAndChild(&Person{name: "Hiroshi"}, &Person{name: "Bob"})
	relationships.relations = append(relationships.relations, Info{
		from:         &Person{name: "Ellie"},
		relationship: Child,
		to:           &Person{name: "Joel"},
	})
	relationships.AddParentAndChild(&Person{name: "John"}, &Person{name: "John Jr"})

	research := ResearchWithInterface{
		browser: &Browser{},
	}

	results, err := research.browser.Investigate(relationships.relations, Child)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Final results: %+v\n", results)
}

func (b *Browser) Investigate(relations []Info, relationType Relationship) ([]Info, error) {
	var matchingRelations []Info
	// relations is accessed in the tight bind to Research
	// e.g. this would error if LLM changes relations to anything but a slice
	for _, relation := range relations {
		// do stuff with relations
		if relation.relationship == relationType {
			matchingRelations = append(matchingRelations, relation)
		}
	}

	// nothing found
	if len(matchingRelations) == 0 {
		return matchingRelations, fmt.Errorf("No matching relations")

	}

	return matchingRelations, nil
}

// MORE EXAMPLES OF DIP

// -- Using Constructor --

type Employee struct {
	id    uint
	name  string
	email string
}

type EmployeeManager struct {
	employees []Employee
}

func (e *EmployeeManager) addPerson(person Employee) []Employee {
	e.employees = append(e.employees, person)
	return e.employees
}

func (e *EmployeeManager) deletePerson(id uint) error {
	updatedEmployees := make([]Employee, len(e.employees))

	found := false

	for _, employee := range e.employees {
		if employee.id == id {
			found = true
		} else {
			updatedEmployees = append(updatedEmployees, employee)
		}
	}

	// still not found, return error
	if !found {
		return fmt.Errorf("Employee with id %d was not found", id)
	}

	e.employees = updatedEmployees
	return nil
}

func NewEmployeeManager() PersonellManager[Employee] {
	return &EmployeeManager{}
}

type PersonellManager[T any] interface {
	addPerson(person T) []T
	deletePerson(id uint) error
}

// -- Without Using Constructor --

// Define an interface for the employee browser.
type EmployeeBrowser interface {
	// InvestigateEmployee provides the flexibility to look for employees by any criteria
	InvestigateEmployee(employeeList []Employee2, id uint) (*Employee2, error)
}

// Employee is our entity type.
type Employee2 struct {
	id    uint
	name  string
	email string
}

// EmployeeRepository implements the EmployeeBrowser interface.
type EmployeeRepository struct {
	employees []Employee2
}

// InvestigateEmployee implements the EmployeeBrowser interface, allowing
// us to investigate the employee list by a specific condition (e.g., by id).
func (r *EmployeeRepository) InvestigateEmployee(employeeList []Employee2, id uint) (*Employee2, error) {
	for _, emp := range employeeList {
		if emp.id == id {
			return &emp, nil
		}
	}
	return nil, fmt.Errorf("Employee with id %d not found", id)
}

// ResearchWithInterface depends on the abstraction (EmployeeBrowser) instead of a concrete implementation.
type ResearchWithInterface2 struct {
	browser EmployeeBrowser
}

// InvestigateEmployeeByID uses the EmployeeBrowser abstraction to search for an employee by id.
func (r *ResearchWithInterface2) InvestigateEmployeeByID(id uint) {
	employee, err := r.browser.InvestigateEmployee(r.browser.(*EmployeeRepository).employees, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Found Employee: %+v\n", employee)
}

func main() {
	// Set up the EmployeeRepository with a few employees.
	repo := &EmployeeRepository{
		employees: []Employee2{
			{id: 1, name: "John", email: "john@example.com"},
			{id: 2, name: "Jane", email: "jane@example.com"},
		},
	}

	// Set up the high-level module ResearchWithInterface which uses the abstraction EmployeeBrowser.
	research := ResearchWithInterface2{browser: repo}

	// Investigate an employee by id using the interface.
	research.InvestigateEmployeeByID(1)
}

/**
* Interface Segregation Principle
**/

package designpatterns

import "fmt"

func Factory() {
	fmt.Println("Factory Design Principle")

}

/**
* Factory Function
**/
type Robot struct {
	name string
	age  uint
}

func NewRobot(name string, age uint) *Robot {
	return &Robot{
		name: name,
		age:  age,
	}
}

type WithIntelligence interface {
	think(string) (string, error)
}

type RobotWithIntelligence struct {
	name string
	age  uint
}

func (r *RobotWithIntelligence) think(thought string) (string, error) {

	if thought == "error" {
		return "", fmt.Errorf("Errored")
	}

	return fmt.Sprintf("Im thinking about %s\n", thought), nil
}

func NewRobotWithIntelligence(name string, age uint) WithIntelligence {
	return &RobotWithIntelligence{
		name: name,
		age:  age,
	}
}

/**
* Factory Generator
**/

type Worker struct {
	name         string
	position     string
	annualIncome int
}

// use a higher order function to create a factory function
func NewWorkerFactory(position string, annualIncome int) func(name string) *Worker {
	return func(name string) *Worker {
		return &Worker{
			name:         name,
			position:     position,
			annualIncome: annualIncome,
		}
	}
}

// Usage
func TestFactoryGenerator() {
	// fixes created workers with this factory function to always be mcdonalds related with
	// the same salary
	mcdonaldsFactory := NewWorkerFactory("mcdonalds cashier", 10000)
	mcdonaldsManagerFactory := NewWorkerFactory("mcdonalds manager", 45000)

	workerBob := mcdonaldsFactory("Bob")
	managerJason := mcdonaldsManagerFactory("Jason")

	fmt.Printf("Worker Bob struct: %+v\n", workerBob)
	fmt.Printf("Manager Jason struct: %+v\n", managerJason)
}

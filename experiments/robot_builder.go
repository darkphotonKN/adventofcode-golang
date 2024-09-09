package experiments

import "fmt"

type Robot struct {
	SerialNo  string
	ModelName string
	Value     int
}

func (r Robot) SetModelNameCopy(name string) string {
	r.ModelName = name
	return r.ModelName
}

func (r *Robot) SetModelName(name string) string {
	r.ModelName = name
	return r.ModelName
}

func NewRobot() *Robot {
	return &Robot{}
}

func RunRobotBuilder() {

	robotCopy := NewRobot()

	// immutable
	robotName := robotCopy.SetModelNameCopy("RobotCopy")
	fmt.Println("RobotCopy name:", robotCopy.ModelName)
	fmt.Println("Copied robot name:", robotName)

	robotRef := NewRobot()
	robotRef.SetModelName("RobotRef")
	fmt.Println("RobotRef name:", robotRef.ModelName)
}

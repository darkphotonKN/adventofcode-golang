package adventofcode2021

import (
	"challenges/tools"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type submarineS struct {
	position int
	depth    int
	aim      int
}

func (sub *submarineS) updatePosition(direction direction, value int) int {

	switch direction {
	case UP:
		sub.depth -= value
		break
	case DOWN:
		sub.depth += value
		break
	case FORWARD:
		sub.position += value
		break
	default:
		log.Println("Direction was invalid")
		return 0
	}

	return 0
}

type direction string

const (
	UP      direction = "up"
	DOWN    direction = "down"
	FORWARD direction = "forward"
)

func Dive() {
	fmt.Println("Dive")
	text := tools.FileReader("./adventofcode/2021/data/dayTwo.txt")

	submarine := submarineS{}

	// loop through instructions and update submarine voyage

	for _, line := range strings.Split(text, "\n") {
		// split into direction and position
		instructions := strings.Split(line, " ")

		fmt.Println("instructions were:", instructions)

		// skip this loop if the instructions are invalid
		if len(instructions) < 2 {
			continue
		}

		position := direction(instructions[0])
		fmt.Println("position was:", position)
		value := instructions[1]
		intValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error when converting value:", err)
		}

		fmt.Println("value was:", intValue)

		submarine.updatePosition(position, intValue)
	}

	log.Println("position:", submarine.position)
	log.Println("depth:", submarine.depth)
	log.Println("position x depth:", submarine.position*submarine.depth)
}

func (s *submarineS) AimModeMove(direction direction, value int) {
	switch direction {
	// only aims upwards, does not move
	case UP:
		s.aim -= value

	// only aims downwards, does not move
	case DOWN:
		s.aim += value

		// moves in the direction of the aim
	case FORWARD:
		s.position += value
		s.depth += value * s.aim
	}
}

func DivePartTwo() {
	fmt.Println("Dive - Part Two")
	text := tools.FileReader("./adventofcode/2021/data/dayTwo.txt")

	instructions := strings.Split(text, "\n")

	submarine := submarineS{}

	for _, instruction := range instructions {
		insSplit := strings.Split(instruction, " ")

		if len(insSplit) < 2 {
			continue
		}

		dir := direction(insSplit[0]) // cast to direction type
		val, err := strconv.Atoi(insSplit[1])

		fmt.Printf("current aim: %d\n", submarine.aim)

		if err != nil {
			fmt.Println("Error when converting value to int: ", err)
			continue
		}

		submarine.AimModeMove(dir, val)
	}

	fmt.Printf("depth: %d\n", submarine.depth)
	fmt.Printf("position: %d\n", submarine.position)
	fmt.Printf("result: %d\n", submarine.depth*submarine.position)
}

/*

-- Part One --

Now, you need to figure out how to pilot this thing.

It seems like the submarine can take a series of commands like forward 1, down 2, or up 3:

forward X increases the horizontal position by X units.
down X increases the depth by X units.
up X decreases the depth by X units.
Note that since you're on a submarine, down and up affect your depth, and so they have the opposite result of what you might expect.

The submarine seems to already have a planned course (your puzzle input). You should probably figure out where it's going. For example:

forward 5
down 5
forward 8
up 3
down 8
forward 2
Your horizontal position and depth both start at 0. The steps above would then modify them as follows:

forward 5 adds 5 to your horizontal position, a total of 5.
down 5 adds 5 to your depth, resulting in a value of 5.
forward 8 adds 8 to your horizontal position, a total of 13.
up 3 decreases your depth by 3, resulting in a value of 2.
down 8 adds 8 to your depth, resulting in a value of 10.
forward 2 adds 2 to your horizontal position, a total of 15.
After following these instructions, you would have a horizontal position of 15 and a depth of 10. (Multiplying these together produces 150.)

Calculate the horizontal position and depth you would have after following the planned course. What do you get if you multiply your final horizontal position by your final depth?

--- Part Two ---
Based on your calculations, the planned course doesn't seem to make any sense. You find the submarine manual and discover that the process is actually slightly more complicated.

In addition to horizontal position and depth, you'll also need to track a third value, aim, which also starts at 0. The commands also mean something entirely different than you first thought:

down X increases your aim by X units.
up X decreases your aim by X units.
forward X does two things:
It increases your horizontal position by X units.
It increases your depth by your aim multiplied by X.
Again note that since you're on a submarine, down and up do the opposite of what you might expect: "down" means aiming in the positive direction.

Now, the above example does something different:

forward 5 adds 5 to your horizontal position, a total of 5. Because your aim is 0, your depth does not change.
down 5 adds 5 to your aim, resulting in a value of 5.
forward 8 adds 8 to your horizontal position, a total of 13. Because your aim is 5, your depth increases by 8*5=40.
up 3 decreases your aim by 3, resulting in a value of 2.
down 8 adds 8 to your aim, resulting in a value of 10.
forward 2 adds 2 to your horizontal position, a total of 15. Because your aim is 10, your depth increases by 2*10=20 to a total of 60.
After following these new instructions, you would have a horizontal position of 15 and a depth of 60. (Multiplying these produces 900.)

Using this new interpretation of the commands, calculate the horizontal position and depth you would have after following the planned course. What do you get if you multiply your final horizontal position by your final depth?
*/

package practice

import (
	"fmt"
	"slices"
)

type User struct {
	name string
	age  int
}

func Run() {
	numbers := []int{20, 7, 10, 3, 4, 3, 14, 7, 20, 3, 15, 4}

	fmt.Println("Sorted list:", sortArrayEvenAndOdd(numbers))

	// updating pointer of slices
	users := []User{
		{
			name: "bob",
			age:  12,
		},
		{
			name: "John",
			age:  33,
		},
	}
	fmt.Printf("\n\nusers BEFORE: %+v\n\n", users)

	// find john
	var foundUser *User

	for _, user := range users {
		if user.name == "John" {
			foundUser = &user
		}
	}

	// update john's age
	foundUser.age = 69

	fmt.Printf("\n\nusers AFTER: %+v\n\n", users)
}

// Sort Function - Sort based on function provided
func sortArrayEvenAndOdd(list []int) []int {
	slices.SortFunc(list, func(a, b int) int {
		// can use a or b, doesn't matter
		if a%2 == 0 {
			return 1
		} else {
			return -1
		}
	})

	return list
}

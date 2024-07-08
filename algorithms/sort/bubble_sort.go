package algorithms

import "fmt"

func BubbleSort(numbers *[]int) {
	fmt.Println("Bubble Sort")

	// dereferencing
	nums := *numbers

	// bubble up the larger number and repeat
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ { // dont need to check the tail
			// compare and do a swap if the number is larger

			// dont need to compare the same index
			if nums[j] > nums[j+1] {
				// swap
				temp := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = temp
			}
		}

		fmt.Println("After one iteration:", numbers)
	}
}

func RunBubbleSort() {
	numbers := []int{22, 9, 13, 4, 50, 19, 2, 1}

	BubbleSort(&numbers)

	fmt.Println("Sorted numbers:", numbers)
}

package utils

import "fmt"

func printNum(num int) {
	fmt.Println("Current number is", num)
}

// Add multiple number together
func Add(nums ...int) int {
	total := 0
	for _, num := range nums {
		printNum(num)
		total += num
	}
	return total
}

package main

import "fmt"

func MaximumSubarraySum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	} else {
		sum := 0
		for _, value := range numbers {
			sum += value
		}

		if sum < 0 {
			return 0
		} else {
			return sum
		}

	}
}

func main() {

	fmt.Println(MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

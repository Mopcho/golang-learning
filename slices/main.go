package main

import "fmt"

func variadicSum(nums ...int) int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}

func main() {
	mySLice := make([]int, 0, 10)

	fmt.Printf("%v", mySLice)

	sum := variadicSum(1, 2, 3)

	fmt.Printf("%v", sum)
}

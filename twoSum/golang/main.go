package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	hasher := make(map[int]int)

	for index, num := range numbers {
		if val, found := hasher[num]; found {
			return []int{val, index}
		}
		hasher[target-num] = index
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3, 12, 4, 7}, 7))
}

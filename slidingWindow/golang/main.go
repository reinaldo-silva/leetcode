package main

import "fmt"

func max(numbers []int) int {

	bigger := numbers[0]
	for _, num := range numbers {
		if num > bigger {
			bigger = num
		}
	}
	return bigger
}

func longestSubArray(numbers []int) int {

	size := len(numbers)
	lo, hi := 0, 0
	curr, _max := 0, 0
	zeroes := 0

	for hi < size {

		if numbers[hi] == 1 {
			curr += 1
			_max = max([]int{_max, curr})
		} else {
			zeroes += 1
		}

		for zeroes > 1 {
			if numbers[lo] == 0 {
				zeroes -= 1
			} else {
				curr -= 1
			}
			lo += 1
		}
		hi += 1

	}

	if _max == size {
		return size - 1
	}

	return _max

}

func main() {
	fmt.Println(longestSubArray([]int{1, 1, 0, 1}))
	fmt.Println(longestSubArray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
	fmt.Println(longestSubArray([]int{1, 1, 1}))
	fmt.Println(longestSubArray([]int{0, 0, 0}))
}

package main

import "fmt"

func main() {
	fmt.Println(LongestSubArray([]int{1, 1, 0, 1}))
	fmt.Println(LongestSubArray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
	fmt.Println(LongestSubArray([]int{1, 1, 1}))
	fmt.Println(LongestSubArray([]int{0, 0, 0}))

	s := "abcdlkasl"
	fmt.Println(LongestSubArrayNoRepeat(s))
	s2 := "abcabcabcabcd"
	fmt.Println(LongestSubArrayNoRepeat(s2))

}

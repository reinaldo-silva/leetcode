package main

func LongestSubArrayNoRepeat(text string) int {

	charSet := make(map[rune]bool)
	left, maxLength := 0, 0

	for right, char := range text {
		for charSet[char] {
			delete(charSet, rune(text[left]))
			left++
		}

		charSet[char] = true
		maxLength = Max([]int{maxLength, right - left + 1})

	}

	return maxLength

}

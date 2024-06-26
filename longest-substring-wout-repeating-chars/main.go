package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	max := make([]int, 0)
	temp := 0
	usedChars := make(map[rune]bool, len(s))

	for i, c := range s {
		_, ok := usedChars[c]

		if !ok {
			temp += 1
			usedChars[c] = true
			fmt.Println("char ", string(c), c, " is not in the map, so temp is", temp)
			fmt.Println(usedChars)
		} else {
			max = append(max, temp)
			temp = 1
			for k := range usedChars {
				delete(usedChars, k)
			}
			usedChars[c] = true
			fmt.Println("char", string(c), "is in the map, so temp is", temp)
			fmt.Println(usedChars)
		}

		if i == len(s)-1 {
			max = append(max, temp)
		}
	}

	return findMax(max)
}

func findMax(l []int) int {
	max := 0

	for _, e := range l {
		if e > max {
			max = e
		}
	}

	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"), "dvdf")
	fmt.Println(lengthOfLongestSubstring("devdf"), "devdf")
	fmt.Println(lengthOfLongestSubstring("pwwkew"), "pwwkew")
	fmt.Println(lengthOfLongestSubstring("aab"), "aab")
}

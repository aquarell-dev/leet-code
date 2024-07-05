package main

import (
	"fmt"
	"sort"
	// "sort"
)

/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.
*/

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string, 0)
	var grouped [][]string

	for _, word := range strs {
		charArray := []rune(word)

		sort.Slice(charArray, func(i int, j int) bool {
			return charArray[i] < charArray[j]
		})

		sorted := string(charArray)

		value, exists := anagrams[sorted]

		if !exists {
			anagrams[sorted] = []string{word}
		} else {
			anagrams[sorted] = append(value, word)
		}
	}

	for _, v := range anagrams {
		grouped = append(grouped, v)
	}

	return grouped
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

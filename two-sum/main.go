package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indexes := make(map[int]int, len(nums))

	for i, num := range nums {
		left := target - num

		value, exists := indexes[left]

		if !exists {
			indexes[num] = i
			continue
		}

		return []int{i, value}
	}

	return make([]int, 2)
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

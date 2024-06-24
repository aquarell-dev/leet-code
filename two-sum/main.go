package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, firstNum := range nums {
		for j, secondNum := range nums {
			if secondNum+firstNum == target && i != j {
				return []int{i, j}
			}
		}
	}
	return make([]int, 2)
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

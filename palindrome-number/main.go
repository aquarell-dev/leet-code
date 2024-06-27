package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	var reversed int
	temp := x

	for {
		if temp == 0 {
			break
		}

		reversed = reversed*10 + (temp % 10)
		temp /= 10
	}

	return reversed == x
}

func main() {
	fmt.Println(isPalindrome(131))
}

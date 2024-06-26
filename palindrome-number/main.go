package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	fmt.Println(math.Floor(math.Log10(float64(x)) + 1))

	return true
}

func main() {
	fmt.Println(isPalindrome(121))
}

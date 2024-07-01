package main

import (
	"fmt"
	"strings"
)

const ASCII_ZERO int = 48

func myAtoi(s string) int {
	var sign, start, converted int

	s = strings.TrimSpace(s)

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		sign = 1
		start = 1
	} else {
		sign = 1
	}

	for _, char := range s[start:] {
		if char < rune(ASCII_ZERO) || char > 57 {
			break
		}

		converted = converted*10 + int(char) - ASCII_ZERO
	}

	converted *= sign

	return converted
}

func main() {
	fmt.Println(myAtoi("  -123"))
}

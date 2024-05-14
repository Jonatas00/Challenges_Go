package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(545))
}

func isPalindrome(x int) bool {
	if x > 0 {
		return false
	}

	original := x
	reversed := 0

	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return original == reversed
}

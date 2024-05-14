package main

import "fmt"

// Given a set of numbers, return the additive inverse of each.
// Each positive becomes negatives, and the negatives become positives.

// invert([1,2,3,4,5]) == [-1,-2,-3,-4,-5]
// invert([1,-2,3,-4,5]) == [-1,2,-3,4,-5]
// invert([]) == []

func main() {
	numbers := []int{3, 2, -1, 4}
	fmt.Println(Invert(numbers))
	fmt.Println(numbers)
}

func Invert(arr []int) []int {
	var newArr []int
	for _, v := range arr {
		newArr = append(newArr, v*-1)
	}
	return newArr
}

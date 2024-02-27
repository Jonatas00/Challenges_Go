package main

import "fmt"

// Given an integer as input, can you round it to the next (meaning, "greater than or equal") multiple of 5?

func main() {
	fmt.Println(RoundToNext5(7))
	fmt.Println(RoundToNext5(2))
	fmt.Println(RoundToNext5(3))
	fmt.Println(RoundToNext5(12))
	fmt.Println(RoundToNext5(-2))
}


func RoundToNext5(n int) int {
	if n >= 0 {
		if n % 5 != 0 {
			return n + (5 - n % 5)
		}
		return n
	} else {
		if n % 5 != 0 {
			return n - (5 + n % 5)
		}
		return n
	}
}

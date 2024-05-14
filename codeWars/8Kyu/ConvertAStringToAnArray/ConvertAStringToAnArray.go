package main

import (
	"fmt"
	"strings"
)

// Write a function to split a string and convert it into an array of words.
// "Robin Singh" ==> ["Robin", "Singh"]
// "I love arrays they are my favorite" ==> ["I", "love", "arrays", "they", "are", "my", "favorite"]

func main() {
	fmt.Println(StringToArray("Robin Singh"))
	fmt.Println(StringToArray("CodeWars"))
	fmt.Println(StringToArray("I love arrays they are my favorite"))
	fmt.Println(StringToArray("1 2 3"))
}

func StringToArray(str string) []string {
	splitedString := strings.Split(str, " ")
	return splitedString
}

package main

import (
	"fmt"
	"strings"
)

//Simple, given a string of words, return the length of the shortest word(s).
//String will never be empty and you do not need to account for different data types.

func main()  {
	fmt.Println(FindShort("bitcoin take over the world maybe who knows perhaps"))
}

func FindShort(s string) int {
	var separatedWord = strings.Split(s, " ")
	var shortWord = separatedWord[0]
	for _, word := range(separatedWord) {
		if len(word) < len(shortWord) {
			shortWord = word
		}
	}

	return len(shortWord)
}

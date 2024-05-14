package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(NoSpace("    dkawodkowa   8988888    dwaokdwoa"))
}

func NoSpace(word string) string {
	newWord := strings.Replace(word, " ", "", -1)
	return newWord
}

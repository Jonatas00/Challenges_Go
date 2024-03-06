package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(ToAlternatingCase("Olá, Mundo!"))
	fmt.Println(ToAlternatingCase("Teste, TestAndo 321 123"))
}

func ToAlternatingCase(str string) string {
	runes := []rune(str)
	for i, v := range runes {
		if unicode.IsUpper(v) {
			runes[i] = unicode.ToLower(v)
		} else if unicode.IsLower(v) {
			runes[i] = unicode.ToUpper(v)
		}
	}

	return string(runes)
}
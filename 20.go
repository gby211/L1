package main

import (
	"fmt"
	"strings"
)

func reverseString2(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		words[i] = reverseString2(word)
	}
	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun"
	reversed := reverseWords(s)
	fmt.Println(reversed) // wons god nus
}

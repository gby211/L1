package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(uniqueChars("abcd"))      //true
	fmt.Println(uniqueChars("abCdefAaf")) //false
	fmt.Println(uniqueChars("aabcd"))     //false

}

func uniqueChars(str string) bool {
	chars := make(map[rune]bool)
	for _, char := range strings.ToLower(str) {
		if chars[char] {
			return false
		}
		chars[char] = true
	}
	return true
}

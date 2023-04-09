package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool)
	for _, s := range sequence {
		set[s] = true
	}

	fmt.Println(set)
}

package main

import "fmt"

func intersection(a, b []int) []int {
	m := make(map[int]bool)

	for _, x := range a {
		m[x] = true
	}

	var res []int

	for _, x := range b {
		if m[x] {
			res = append(res, x)
		}
	}

	return res
}

func main() {
	a := []int{1, 2, 3}
	b := []int{3, 5, 1}

	fmt.Println(intersection(a, b)) // [3 1]
}

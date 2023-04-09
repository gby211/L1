package main

import (
	"fmt"
	"sort"
)

//sort встроенная функция реализующая быструю сортировку

// Slice sorts the slice x given the provided less function.
// It panics if x is not a slice.

// pdqsort_func sorts data[a:b].
// The algorithm based on pattern-defeating quicksort(pdqsort), but without the optimizations from BlockQuicksort.
// pdqsort paper: https://arxiv.org/pdf/2106.05123.pdf

func main() {
	arr := []int{5, 1, 7, 3, 9, 2, 8, 4, 6}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	fmt.Println(arr)
}

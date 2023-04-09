package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	wg.Add(len(numbers))

	sum := 0

	for _, num := range numbers {
		go func(n int) {
			defer wg.Done()
			square := n * n
			sum += square
		}(num)
	}

	wg.Wait()

	fmt.Print(sum)
}

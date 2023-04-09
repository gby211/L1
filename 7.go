package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	myMap := make(map[int]int)

	for i := 0; i < 10; i++ {
		go func(val int) {
			mu.Lock()
			myMap[val] = val
			mu.Unlock()
		}(i)
	}

	for len(myMap) < 10 {
	}

	fmt.Println(myMap)
}

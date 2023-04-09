package main

import (
	"fmt"
	"math"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempGroups := make(map[int][]float64)

	for _, temp := range temps {
		groupKey := int(math.Round(temp/10.0)) * 10
		tempGroups[groupKey] = append(tempGroups[groupKey], temp)
	}

	for key, group := range tempGroups {
		fmt.Printf("%d: %v\n", key, group)
	}
}

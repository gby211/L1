package main

import (
	"fmt"
	"sort"
)

//sort встроенная функция реализующая бинарный поиск

func main() {
	arr := []int{1, 2, 3, 4}

	el := 3

	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= el
	})

	if index < len(arr) && arr[index] == el {
		fmt.Println("Найдено")
	} else {
		fmt.Println("Не найдено")
	}
}

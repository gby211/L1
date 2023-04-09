package main

import (
	"fmt"
)

func main() {
	inputCh := make(chan int)
	outputCh := make(chan int)

	// горутина для чтения данных из входного канала
	go func() {
		for _, x := range []int{1, 2, 3, 4, 5} {
			inputCh <- x
		}
		close(inputCh)
	}()

	// горутина для выполнения операции умножения на 2
	go func() {
		for x := range inputCh {
			outputCh <- x * 2
		}
		close(outputCh)
	}()

	// горутина для вывода результатов в stdout
	go func() {
		for x := range outputCh {
			fmt.Println(x)
		}
	}()

	// ожидание завершения программы по нажатию Ctrl+C
	select {}
}

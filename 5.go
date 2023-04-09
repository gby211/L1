package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; ; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	timeout := 5 * time.Second
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	for {
		select {
		case val := <-ch:
			fmt.Printf("Received value: %d\n", val)
		case <-timer.C:
			fmt.Println("Time is up!")
			return
		}
	}
}

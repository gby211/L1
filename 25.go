package main

import (
	"fmt"
	"time"
)

func mySleep(duration time.Duration) {
	select {
	case <-time.After(duration):
		return
	}
}

func main() {
	fmt.Println("Go")
	mySleep(time.Second)
	fmt.Println("End")
}

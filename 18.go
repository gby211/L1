package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := Counter{value: 0}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter value:", counter.Value())
}

type Counter struct {
	value int
	mux   sync.Mutex
}

func (c *Counter) Inc() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.value
}

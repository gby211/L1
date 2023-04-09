package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Printf("Worker %d: channel closed\n", id)
				return
			}
			fmt.Printf("Worker %d: received %d\n", id, val)
		}
	}
}

func main() {
	// Создаем канал и воркеров
	ch := make(chan int)
	numWorkers := 5
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(i, ch, &wg)
	}

	// Записываем данные в канал
	i := 1
	go func() {
		for {
			ch <- i
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Обработка сигналов завершения программы
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan
	close(ch)
	wg.Wait()
	fmt.Println("All workers have finished.")
}

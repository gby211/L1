package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	third()
}

//С помощью пакета context. Пакет context предоставляет механизм для отмены операции через иерархию вызовов. Для этого
//используется функция context.WithCancel, которая возвращает контекст и функцию cancel.
//При вызове функции cancel, все дочерние контексты и функции, связанные с этим контекстом, будут отменены.

func first() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Cancelled")
				return
			default:
				fmt.Println("Working")
			}
		}
	}()
	time.Sleep(time.Millisecond)
	cancel()
	time.Sleep(time.Millisecond)
	fmt.Println("Exiting...")
}

//Использование канала для остановки горутины.
//Горутина будет выполнять свою работу, пока не получит сообщение через канал.

func second() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Stopped")
				return
			default:
				fmt.Println("Working")
			}
		}
	}()
	time.Sleep(time.Millisecond)
	stop <- true
	time.Sleep(time.Millisecond)
	fmt.Println("Exiting...")
}

// Завершение функции, которая запускается в горутине
func third() {
	go func() {
		fmt.Println("Goroutine is exiting")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Main function is exiting")
}

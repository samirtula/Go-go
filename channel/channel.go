package main

import "fmt"

func main() {
	numCh := make(chan int)
	// Go, запись в канал блокируется до тех пор, пока кто-то не прочитает из этого канала, и наоборот.
	// Это вызывает взаимную блокировку (deadlock)
	go func() {
		numCh <- 10
	}()
	num := <-numCh
	fmt.Println(num)
}

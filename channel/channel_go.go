package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(maxSum([]int{1, 2, 3}, []int{10, 20, 30}))

	msgCh := make(chan string)
	// вызываем функцию асинхронно в горутине
	go printer(msgCh)

	msgCh <- "start"
	msgCh <- "still work"
	msgCh <- "end"

	close(msgCh)
	// ждем, пока printer закончит работу без этого не получим Printed
	time.Sleep(100 * time.Millisecond)
}

// суммирует значения каждого слайса nums и возвращает тот, который имеет наибольшую сумму
func maxSum(a []int, b []int) []int {
	s1Ch := make(chan int)
	go sumParallel(a, s1Ch)

	s2Ch := make(chan int)
	go sumParallel(b, s2Ch)

	s1, s2 := <-s1Ch, <-s2Ch

	if s1 > s2 {
		return a
	}

	return b
}

func sumParallel(nums []int, resCh chan int) {
	s := 0
	for _, n := range nums {
		s += n
	}

	resCh <- s
}

func printer(msgCh chan string) {
	// читаем из канала, пока он открыт
	for msg := range msgCh {
		fmt.Println(msg)
	}

	fmt.Println("Printed")
}

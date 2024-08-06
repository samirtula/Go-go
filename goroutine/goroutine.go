package main

import (
	"fmt"
	"time"
)

func main() {
	//for i := 0; i < 5; i++ {
	//	go func() {
	//		fmt.Println(i)
	//}()
	//}

	//time.Sleep(100 * time.Millisecond)

	// Горутины выполняются независимо и не гарантируют порядка.
	// При необходимости последовательность в выполнении придется реализовывать самостоятельно.
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
}

package main

import (
	"sync"
	"fmt"
)

func externalHTTPNum() int {
	return 1
}

func main() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	sum := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			n := externalHTTPNum()

			mu.Lock()
			sum += n
			mu.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(sum)
}

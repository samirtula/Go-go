package main

import "fmt"

// NumCount alias для существующего типа
type NumCount int

type ErrorCode string

type counter int

// метод типа, напоминает структуру
func (c *counter) inc() {
	*c++
}

func main() {
	n := NumCount(len([]int{1, 2, 3}))
	fmt.Println(n)

	ec := ErrorCode("iternal")
	fmt.Println(string(ec)) //

	cn := counter(5)
	(&cn).inc()
	fmt.Println(cn)
}

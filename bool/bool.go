package main

import "fmt"

func main() {
	fmt.Println(false == false)
	fmt.Println(true && false)
	fmt.Println(false || true)
	fmt.Println(false == false)

	b := true
	fmt.Println(b)
}
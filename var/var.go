package main

import (
	"fmt"
	"errors"
)

var num int = 11

func main() {
	num := 33
	num = 44
	fmt.Println(num)

	longStringName := "StringName"
	fmt.Println(longStringName)

	vars()
}


func vars()  {
	var (
		a string
		b bool
		c int
		errCannotSum = errors.New("cannot sum")
	)

	fmt.Println(a, b, c, errCannotSum)
}